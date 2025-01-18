package codegenerator

import (
	"errors"
	"fmt"
	"github.com/vainback/six-util/v3"
	"os"
	"six-go/database/db"
	"six-go/database/models"
	"six-go/utils"
	"slices"
	"strings"
)

// todo 所有路径不允许带最后的斜线
const (
	// AutoWriteToProject 是否自动写入项目目录
	AutoWriteToProject = true
	// ProjectFrontRootPath 项目前端根目录路径
	ProjectFrontRootPath = "D:\\coding\\learning\\open-resources\\six-admin\\six-arco"
	// ProjectServerRootPath 项目后端根目录路径
	ProjectServerRootPath = "D:\\coding\\learning\\open-resources\\six-admin\\six-go"
	// DownloadFilepath 当选择不自动写入项目目录的时候 点击下载代码自动下载到的位置 会在该目录后面 自动建一层 CodeGenerator 目录
	DownloadFilepath = "D:\\coding\\learning\\open-resources\\"
)

// 字段类型 int(11) bigint(20) varchar(255) text(textarea) float decimal(10,2)-高精度浮点数
// 下拉框 下拉框多选 单选框 复选框 开关(bool) 上传图片 上传图片（多图）上传文件 时间选择器 日期选择器 日期时间选择器 富文本

type CodeGenerator struct {
	IsTree     bool
	v          models.CodeGenerator
	err        error
	auth       string
	parentRule models.AuthRule
}

func Generator(v models.CodeGenerator) error {
	if six.XStr(v.Table).IsEmpty() {
		return errors.New("表名不能为空")
	}

	if v.Fields == nil || len(*v.Fields) == 0 {
		return errors.New("表字段不能为空")
	}

	if six.XStr(v.Title).IsEmpty() {
		return errors.New("菜单标题不能为空")
	}

	if v.ParentModule < 1 {
		return errors.New("上级菜单不能为空，不建议将模块放置在顶级菜单，请先建好上级菜单")
	}

	var parentRule models.AuthRule
	if err := db.DB().Where("id = ?", v.ParentModule).First(&parentRule).Error; err != nil {
		return err
	}

	var isTreeTable bool
	for _, field := range *v.Fields {
		if six.Str(field.Name).TrimSpace() == "" {
			return fmt.Errorf("字段名不能为空: %+v", field)
		}

		if six.Str(field.Type).TrimSpace() == "" {
			return fmt.Errorf("字段类型不能为空: %+v", field)
		}

		if six.Str(field.Name).TrimSpace() == "parent_id" {
			isTreeTable = true
		}

		if field.Default != nil && *field.Default == "null" {
			field.Default = nil
		}
	}

	var g = new(CodeGenerator)
	g.IsTree = isTreeTable
	g.v = v
	g.auth = six.Strings(parentRule.Auth, ":", six.Str(v.Table).Lower().String()).String()
	return g.model().controller().route().api().index().form().finished()
}

func (c *CodeGenerator) model() *CodeGenerator {
	if c.err != nil {
		return c
	}
	tplFilepath := six.Str(ProjectServerRootPath).Append("/static/generator_tpl/server/model.go.tpl").String()
	file, err := os.ReadFile(tplFilepath)
	if err != nil {
		c.err = err
		return c
	}

	sqlFile, err := os.ReadFile(six.Str(ProjectServerRootPath).Append("/static/generator_tpl/server/create_table.sql.tpl").String())
	if err != nil {
		c.err = err
		return c
	}

	createTableSql := string(sqlFile)
	createTableSql = strings.ReplaceAll(createTableSql, "{{tpl - $table}}", strings.TrimSpace(c.v.Table))
	createTableSql = strings.Replace(createTableSql, "<tpl - $delete_time>", utils.Ternary(c.v.IsSoftDelete, "`delete_time` datetime(3) NULL DEFAULT NULL,", ""), 1)
	createTableSql = strings.Replace(createTableSql, "<tpl - $tenant>", utils.Ternary(c.v.IsTenant, "`tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',", ""), 1)

	content := string(file)
	content = strings.ReplaceAll(content, "<Tpl-Module>", six.Str(c.v.Table).Snake2BigCamel())
	content = strings.ReplaceAll(content, "<Tpl-TableName>", six.Str(c.v.Table).TrimSpace())
	content = strings.ReplaceAll(content, "<Tpl-SoftDelete>", utils.Ternary(c.v.IsSoftDelete, "db.SoftDelete", "<Tpl-Space-Line>"))
	content = strings.ReplaceAll(content, "<Tpl-Tenant>", utils.Ternary(c.v.IsTenant, "TenantIdentify", "<Tpl-Space-Line>"))
	content = strings.ReplaceAll(content, "<Tpl-TenantWhere>", utils.Ternary(c.v.IsTenant, ".Scopes(data.TenantIdentify.SqlBuilder(data.TableName()))", ""))

	var fields strings.Builder
	var requireds strings.Builder
	var keywords strings.Builder
	var sqlFields []string
	var sqlIndexes []string

	if c.v.IsSoftDelete {
		sqlIndexes = append(sqlIndexes, "INDEX `idx_auth_user_delete_time`(`delete_time`) USING BTREE,")
	}
	if c.v.IsTenant {
		sqlIndexes = append(sqlIndexes, "INDEX `idx_auth_relation_tenant_id`(`tenant_id`) USING BTREE,")
	}

	tpl := "`json:\"{{tpl - $name}}\" gorm:\"type:{{tpl - $type}};column:{{tpl - $name}};{{tpl - $comment}}{{tpl - $default}}{{tpl - $index}}\"`"

	var hasTime bool
	for _, f := range *c.v.Fields {
		f.Type = strings.TrimSpace(f.Type)
		f.Name = strings.TrimSpace(f.Name)
		f.Title = strings.TrimSpace(f.Title)
		f.Comment = strings.TrimSpace(f.Comment)
		f.Index = strings.TrimSpace(f.Index)

		tag := strings.ReplaceAll(tpl, "{{tpl - $name}}", f.Name)
		var words []string
		words = append(words, six.Str(f.Name).Snake2BigCamel())

		isDefault := f.Default != nil
		var defaultValue string
		if isDefault {
			defaultValue = *f.Default
		}

		if strings.TrimSpace(f.Index) == "index" {
			sqlIndexes = append(sqlIndexes, fmt.Sprintf("INDEX `idx_%s_%s`(`%s`) USING BTREE,", strings.TrimSpace(c.v.Table), f.Name, f.Name))
		}
		if strings.TrimSpace(f.Index) == "unique" {
			sqlIndexes = append(sqlIndexes, fmt.Sprintf("UNIQUE INDEX `uni_%s_%s`(`%s`) USING BTREE,", strings.TrimSpace(c.v.Table), f.Name, f.Name))
		}

		if slices.Contains([]string{"int(11)", "bigint(20)"}, f.Type) {
			tag = strings.Replace(tag, "{{tpl - $type}}", f.Type, 1)
			tag = strings.Replace(tag, "{{tpl - $comment}}", utils.Ternary(strings.TrimSpace(f.Comment) != "", fmt.Sprintf("comment:%s;", f.Comment), ""), 1)
			tag = strings.Replace(tag, "{{tpl - $default}}", utils.Ternary(f.Default != nil, fmt.Sprintf("not null;default:%d;", six.Str(defaultValue).Int()), ""), 1)
			tag = strings.Replace(tag, "{{tpl - $index}}", utils.Ternary(strings.TrimSpace(f.Index) != "", f.Index+";", ""), 1)

			sqlFields = append(sqlFields, fmt.Sprintf("`%s` %s %s COMMENT '%s',", f.Name, f.Type, utils.Ternary(f.Default != nil, fmt.Sprintf("not null default %d", six.Str(defaultValue).Int()), ""), f.Comment))
			words = append(words, "int64", tag)
		} else if strings.Contains(f.Type, "varchar") {
			tag = strings.Replace(tag, "{{tpl - $type}}", f.Type, 1)
			tag = strings.Replace(tag, "{{tpl - $comment}}", utils.Ternary(strings.TrimSpace(f.Comment) != "", fmt.Sprintf("comment:%s;", f.Comment), ""), 1)
			tag = strings.Replace(tag, "{{tpl - $default}}", utils.Ternary(f.Default != nil, fmt.Sprintf("not null;default:'%s';", defaultValue), ""), 1)
			tag = strings.Replace(tag, "{{tpl - $index}}", utils.Ternary(strings.TrimSpace(f.Index) != "", f.Index+";", ""), 1)

			sqlFields = append(sqlFields, fmt.Sprintf("`%s` %s CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci %s COMMENT '%s',", f.Name, f.Type, utils.Ternary(f.Default != nil, fmt.Sprintf("not null default '%s'", defaultValue), ""), f.Comment))

			words = append(words, "string", tag)
		} else if strings.Contains(f.Type, "text") {
			tag = strings.Replace(tag, "{{tpl - $type}}", f.Type, 1)
			tag = strings.Replace(tag, "{{tpl - $comment}}", utils.Ternary(strings.TrimSpace(f.Comment) != "", fmt.Sprintf("comment:%s;", f.Comment), ""), 1)
			tag = strings.Replace(tag, "{{tpl - $index}}", utils.Ternary(strings.TrimSpace(f.Index) != "", f.Index+";", ""), 1)

			sqlFields = append(sqlFields, fmt.Sprintf("`%s` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '%s',", f.Name, f.Comment))

			words = append(words, "string", tag)
		} else if strings.Contains(f.Type, "float") || strings.Contains(f.Type, "decimal") {
			tag = strings.Replace(tag, "{{tpl - $type}}", f.Type, 1)
			tag = strings.Replace(tag, "{{tpl - $comment}}", utils.Ternary(strings.TrimSpace(f.Comment) != "", fmt.Sprintf("comment:%s;", f.Comment), ""), 1)
			tag = strings.Replace(tag, "{{tpl - $default}}", utils.Ternary(f.Default != nil, fmt.Sprintf("not null;default:%.2f;", six.Str(defaultValue).Float()), ""), 1)
			tag = strings.Replace(tag, "{{tpl - $index}}", utils.Ternary(strings.TrimSpace(f.Index) != "", f.Index+";", ""), 1)

			sqlFields = append(sqlFields, fmt.Sprintf("`%s` %s %s COMMENT '%s',", f.Name, f.Type, utils.Ternary(f.Default != nil, fmt.Sprintf("not null default %.2f", six.Str(defaultValue).Float()), ""), f.Comment))

			words = append(words, "float64", tag)
		} else if f.Type == "下拉框" || f.Type == "单选框" {
			tag = strings.Replace(tag, "{{tpl - $type}}", "varchar(32)", 1)
			tag = strings.Replace(tag, "{{tpl - $comment}}", utils.Ternary(strings.TrimSpace(f.Comment) != "", fmt.Sprintf("comment:%s;", f.Comment), ""), 1)
			tag = strings.Replace(tag, "{{tpl - $default}}", utils.Ternary(f.Default != nil, fmt.Sprintf("not null;default:'%s';", defaultValue), ""), 1)
			tag = strings.Replace(tag, "{{tpl - $index}}", utils.Ternary(strings.TrimSpace(f.Index) != "", f.Index+";", ""), 1)

			sqlFields = append(sqlFields, fmt.Sprintf("`%s` %s CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci %s COMMENT '%s',", f.Name, "varchar(32)", utils.Ternary(f.Default != nil, fmt.Sprintf("not null default '%s'", defaultValue), ""), f.Comment))

			words = append(words, "string", tag)
		} else if f.Type == "下拉框多选" || f.Type == "复选框" {
			tag = strings.Replace(tag, "{{tpl - $type}}", "text;serializer:json", 1)
			tag = strings.Replace(tag, "{{tpl - $comment}}", utils.Ternary(strings.TrimSpace(f.Comment) != "", fmt.Sprintf("comment:%s;", f.Comment), ""), 1)
			tag = strings.Replace(tag, "{{tpl - $index}}", utils.Ternary(strings.TrimSpace(f.Index) != "", f.Index+";", ""), 1)

			sqlFields = append(sqlFields, fmt.Sprintf("`%s` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '%s',", f.Name, f.Comment))

			words = append(words, "[]string", tag)
		} else if f.Type == "开关" {
			tag = strings.Replace(tag, "{{tpl - $type}}", "tinyint(1)", 1)
			tag = strings.Replace(tag, "{{tpl - $comment}}", utils.Ternary(strings.TrimSpace(f.Comment) != "", fmt.Sprintf("comment:%s;", f.Comment), ""), 1)
			tag = strings.Replace(tag, "{{tpl - $default}}", utils.Ternary(f.Default != nil, fmt.Sprintf("not null;default:%d;", six.XBool(six.Str(defaultValue).Bool()).Int()), ""), 1)
			tag = strings.Replace(tag, "{{tpl - $index}}", utils.Ternary(strings.TrimSpace(f.Index) != "", f.Index+";", ""), 1)

			sqlFields = append(sqlFields, fmt.Sprintf("`%s` %s %s COMMENT '%s',", f.Name, "tinyint(1)", utils.Ternary(f.Default != nil, fmt.Sprintf("not null default %d", six.XBool(six.Str(defaultValue).Bool()).Int()), ""), f.Comment))

			words = append(words, "bool", tag)
		} else if f.Type == "上传图片" || f.Type == "上传图片（多图）" || f.Type == "上传文件" || f.Type == "上传视频" {
			tag = strings.Replace(tag, "{{tpl - $type}}", "text;serializer:json", 1)
			tag = strings.Replace(tag, "{{tpl - $comment}}", utils.Ternary(strings.TrimSpace(f.Comment) != "", fmt.Sprintf("comment:%s;", f.Comment), ""), 1)

			sqlFields = append(sqlFields, fmt.Sprintf("`%s` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '%s',", f.Name, f.Comment))

			words = append(words, "[]string", tag)
		} else if f.Type == "时间选择器" || f.Type == "日期选择器" || f.Type == "日期时间选择器" {
			hasTime = true
			tag = strings.Replace(tag, "{{tpl - $type}}", "varchar(32)", 1)
			tag = strings.Replace(tag, "{{tpl - $comment}}", utils.Ternary(strings.TrimSpace(f.Comment) != "", fmt.Sprintf("comment:%s;", f.Comment), ""), 1)
			tag = strings.Replace(tag, "{{tpl - $index}}", utils.Ternary(strings.TrimSpace(f.Index) != "", f.Index+";", ""), 1)

			sqlFields = append(sqlFields, fmt.Sprintf("`%s` %s CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci %s COMMENT '%s',", f.Name, "varchar(32)", utils.Ternary(f.Default != nil, fmt.Sprintf("not null default '%s'", defaultValue), ""), f.Comment))
			words = append(words, "string", tag)
		} else if f.Type == "富文本" {
			tag = strings.Replace(tag, "{{tpl - $type}}", "longtext", 1)
			tag = strings.Replace(tag, "{{tpl - $comment}}", utils.Ternary(strings.TrimSpace(f.Comment) != "", fmt.Sprintf("comment:%s;", f.Comment), ""), 1)
			tag = strings.Replace(tag, "{{tpl - $index}}", utils.Ternary(strings.TrimSpace(f.Index) != "", f.Index+";", ""), 1)

			sqlFields = append(sqlFields, fmt.Sprintf("`%s` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '%s',", f.Name, f.Comment))

			words = append(words, "string", tag)
		}

		fields.WriteString("    " + strings.Join(words, " ") + "\n")
		if f.IsRequired {
			requireds.WriteString(fmt.Sprintf("        validation.Field(&data.%s, validation.Required.Error(\"%s不能为空\")),\n", six.Str(f.Name).Snake2BigCamel(), f.Title))
		}

		if f.IsKeyword {
			keywords.WriteString(fmt.Sprintf("        data.HasTableName(\"%s\"),\n", six.Str(f.Name).TrimSpace()))
		}
	}

	content = strings.Replace(content, "<Tpl - Time - Import>", utils.Ternary(hasTime, `"time"`, ""), 1)
	content = strings.ReplaceAll(content, "<Tpl-Fields>", strings.TrimSpace(fields.String()))
	content = strings.ReplaceAll(content, "{{tpl - $default}}", "")

	if requireds.Len() > 0 {
		content = strings.ReplaceAll(content, "<Tpl-ValidationRequired>", fmt.Sprintf("validation.ValidateStruct(&data,\n%s\n    )", strings.Trim(requireds.String(), "\n")))
	} else {
		content = strings.ReplaceAll(content, "<Tpl-ValidationRequired>", "nil")
	}

	if keywords.Len() > 0 {
		content = strings.ReplaceAll(content, "<Tpl-KeywordsFields>", fmt.Sprintf("six.Arrays(\n%s\n    )", "        "+strings.TrimSpace(keywords.String())))
	} else {
		content = strings.ReplaceAll(content, "<Tpl-KeywordsFields>", "[]string{}")
	}

	split := strings.Split(content, "\n")
	var endStr strings.Builder
	for _, line := range split {
		if !strings.Contains(line, "<Tpl-Space-Line>") {
			endStr.WriteString(line)
		}
	}

	if c.IsTree {
		content = strings.Replace(content, "<Tpl - HasChildren>", "", -1)
	} else {
		content = content[:strings.Index(content, "<Tpl - HasChildren>")]
	}

	createTableSql = strings.Replace(createTableSql, "<tpl - $fields>", strings.Join(sqlFields, "\n  "), 1)
	createTableSql = strings.Replace(createTableSql, "<tpl - $fields-index>", strings.Join(sqlIndexes, "\n  "), 1)

	var execSql []string
	for _, s := range strings.Split(createTableSql, "\n") {
		if strings.TrimSpace(s) != "" {
			execSql = append(execSql, s)
		}
	}

	if err = db.DB().Exec("DROP TABLE IF EXISTS `" + strings.TrimSpace(c.v.Table) + "`").Error; err != nil {
		c.err = err
		return c
	}

	lastSql := strings.Join(execSql, " ")
	lastDou := strings.LastIndex(lastSql, ",")
	lastSql = lastSql[:lastDou] + lastSql[lastDou+1:]
	if err = db.DB().Exec(lastSql).Error; err != nil {
		c.err = err
		return c
	}

	wirtePath := six.Str(ProjectServerRootPath).Append("/database/models/", six.Str(c.v.Table).RemoveSpace(), ".go").String()

	c.err = os.WriteFile(wirtePath, []byte(content), 0666)
	return c
}

func (c *CodeGenerator) controller() *CodeGenerator {
	if c.err != nil {
		return c
	}

	tplFilepath := six.Str(ProjectServerRootPath).Append("/static/generator_tpl/server/controller.go.tpl").String()
	if c.IsTree {
		tplFilepath = strings.ReplaceAll(tplFilepath, "controller.go.tpl", "controller.tree.go.tpl")
	}
	file, err := os.ReadFile(tplFilepath)
	if err != nil {
		c.err = err
		return c
	}

	content := string(file)
	content = strings.ReplaceAll(content, "<Tpl-Module>", six.Str(c.v.Table).Snake2BigCamel())
	wirtePath := six.Str(ProjectServerRootPath).Append("/app/admin/controller/", six.Str(c.v.Table).RemoveSpace(), ".go").String()

	c.err = os.WriteFile(wirtePath, []byte(content), 0666)
	return c
}

func (c *CodeGenerator) route() *CodeGenerator {
	if c.err != nil {
		return c
	}

	tplFilepath := six.Str(ProjectServerRootPath).Append("/static/generator_tpl/server/route.go.tpl").String()
	file, err := os.ReadFile(tplFilepath)
	if err != nil {
		c.err = err
		return c
	}

	content := string(file)

	if !c.IsTree {
		content = strings.ReplaceAll(content, "\n    p.POST(\"/tree-select\", controller.<Tpl-Module>.TreeSelect)", "")
	}

	content = strings.ReplaceAll(content, "<Tpl-Module>", six.Str(c.v.Table).Snake2BigCamel())
	content = strings.ReplaceAll(content, "<Tpl-Route-Group>", six.Str(c.v.Table).Lower().String())

	wirtePath := six.Str(ProjectServerRootPath).Append("/app/admin/route/", six.Str(c.v.Table).RemoveSpace(), ".go").String()

	c.err = os.WriteFile(wirtePath, []byte(content), 0666)

	return c
}

func (c *CodeGenerator) api() *CodeGenerator {
	if c.err != nil {
		return c
	}
	tplFilepath := six.Str(ProjectServerRootPath).Append("/static/generator_tpl/front/api.ts.tpl").String()
	file, err := os.ReadFile(tplFilepath)
	if err != nil {
		c.err = err
		return c
	}

	content := string(file)
	content = strings.ReplaceAll(content, "<Tpl-Module>", six.Str(c.v.Table).Snake2BigCamel())
	content = strings.ReplaceAll(content, "<Tpl-Route-Group>", six.Str(c.v.Table).Lower().String())

	var fields strings.Builder
	var defaultValue strings.Builder

	for _, f := range *c.v.Fields {
		if f.Default != nil {
			if strings.Contains(f.Type, "int") ||
				strings.Contains(f.Type, "float") ||
				strings.Contains(f.Type, "decimal") {

				fields.WriteString(fmt.Sprintf("    %s: number;\n", f.Name))
				defaultValue.WriteString(fmt.Sprintf("    %s: 0,\n", f.Name))

			} else if strings.Contains(f.Type, "varchar") ||
				strings.Contains(f.Type, "text") ||
				strings.Contains(f.Type, "富文本") ||
				f.Type == "下拉框" ||
				strings.Contains(f.Type, "单选框") {

				fields.WriteString(fmt.Sprintf("    %s: string;\n", f.Name))
				defaultValue.WriteString(fmt.Sprintf("    %s: '',\n", f.Name))
			} else if strings.Contains(f.Type, "时间选择器") ||
				strings.Contains(f.Type, "日期选择器") ||
				strings.Contains(f.Type, "日期时间选择器") {
				fields.WriteString(fmt.Sprintf("    %s: string | null;\n", f.Name))
				defaultValue.WriteString(fmt.Sprintf("    %s: null,\n", f.Name))
			} else if strings.Contains(f.Type, "下拉框多选") || strings.Contains(f.Type, "复选框") || strings.Contains(f.Type, "上传") {

				fields.WriteString(fmt.Sprintf("    %s: string[];\n", f.Name))
				defaultValue.WriteString(fmt.Sprintf("    %s: [],\n", f.Name))

			} else if strings.Contains(f.Type, "开关") {
				fields.WriteString(fmt.Sprintf("    %s: boolean;\n", f.Name))
				defaultValue.WriteString(fmt.Sprintf("    %s: false,\n", f.Name))
			}
		} else {
			if strings.Contains(f.Type, "int") ||
				strings.Contains(f.Type, "float") ||
				strings.Contains(f.Type, "decimal") {

				fields.WriteString(fmt.Sprintf("    %s: number | null;\n", f.Name))
				defaultValue.WriteString(fmt.Sprintf("    %s: null,\n", f.Name))

			} else if strings.Contains(f.Type, "varchar") ||
				strings.Contains(f.Type, "text") ||
				strings.Contains(f.Type, "富文本") ||
				strings.TrimSpace(f.Type) == "下拉框" ||
				strings.Contains(f.Type, "单选框") ||
				strings.Contains(f.Type, "时间选择器") ||
				strings.Contains(f.Type, "日期选择器") ||
				strings.Contains(f.Type, "日期时间选择器") {

				fields.WriteString(fmt.Sprintf("    %s: string | null;\n", f.Name))
				defaultValue.WriteString(fmt.Sprintf("    %s: null,\n", f.Name))

			} else if strings.Contains(f.Type, "下拉框多选") || strings.Contains(f.Type, "复选框") || strings.Contains(f.Type, "上传") {

				fields.WriteString(fmt.Sprintf("    %s: string[];\n", f.Name))
				defaultValue.WriteString(fmt.Sprintf("    %s: [],\n", f.Name))

			} else if strings.Contains(f.Type, "开关") {
				fields.WriteString(fmt.Sprintf("    %s: bool;\n", f.Name))
				defaultValue.WriteString(fmt.Sprintf("    %s: false,\n", f.Name))
			}
		}
	}

	content = strings.ReplaceAll(content, "<Tpl-FieldsType>", strings.TrimSpace(fields.String()))
	content = strings.ReplaceAll(content, "<Tpl-FieldsDefaultValue>", strings.TrimSpace(defaultValue.String()))

	wirtePath := six.Str(ProjectFrontRootPath).Append("/src/api/codegen/", six.Str(c.v.Table).RemoveSpace(), ".ts").String()

	c.err = os.WriteFile(wirtePath, []byte(content), 0666)
	return c
}

func (c *CodeGenerator) index() *CodeGenerator {
	if c.err != nil {
		return c
	}
	tplFilepath := six.Str(ProjectServerRootPath).Append("/static/generator_tpl/front/index.vue.tpl").String()
	if c.IsTree {
		tplFilepath = strings.ReplaceAll(tplFilepath, "index.vue.tpl", "index.tree.vue.tpl")
	}
	file, err := os.ReadFile(tplFilepath)
	if err != nil {
		c.err = err
		return c
	}

	content := string(file)
	content = strings.ReplaceAll(content, "<Tpl-Module>", six.Str(c.v.Table).Snake2BigCamel())
	content = strings.ReplaceAll(content, "<Tpl-Module-Lower>", six.Str(c.v.Table).RemoveSpace())
	content = strings.ReplaceAll(content, "<Tpl-AuthName>", c.auth)

	switchTpl := `
<a-table-column title="{{tpl - $title}}" data-index="{{tpl - $name}}">
	<template #cell="{ record }">
		<a-tag v-if="record.status" color="blue"
			>正常
		</a-tag>
		<a-tag v-else>禁用</a-tag>
	</template>
</a-table-column>`

	templateTpl := `
<a-table-column title="{{tpl - $title}}" data-index="{{tpl - $name}}">
	<template #cell="{ record }">
		{{tpl - $tags}}
	</template>
</a-table-column>`

	dictRequest := `
    const dictMap = ref({})
	const dictList = ref({})
    const requestDict = async () => {
        const { data } = await reqDict('select', defaultRequestParam({...EmptyDict}));
        let m = {}
		let l = {}
        for (const key in data) {
            let item = data[key]
            if (!m[item.type]) {
                m[item.type] = {}
            }
			if (!l[item.type]) {
				l[item.type] = []
			}
            m[item.type][item.value] = item
			l[item.type].push(item)
        }
        dictMap.value = {...m}
		dictList.value = {...l}
    }
`

	var columns strings.Builder
	var hasDict bool
	for _, f := range *c.v.Fields {
		f.Type = strings.TrimSpace(f.Type)
		if f.Type == "下拉框" || f.Type == "单选框" {
			content = strings.Replace(content, "<Tpl - Import - Dict>", "import { Dict, reqDict, EmptyDict } from '@/api/basic/dict';", 1)
			content = strings.Replace(content, "<Tpl - Import - Dict - Init>", "requestDict();", 1)
			content = strings.Replace(content, "<Tpl - Request - Dict - Object>", dictRequest, 1)

			kks := fmt.Sprintf("dictMap['%s'][record.%s]", strings.TrimSpace(f.DictType), strings.TrimSpace(f.Name))
			tag := fmt.Sprintf("<a-tag :color=\"%s.color\">{{%s.label}}</a-tag>", kks, kks)

			item := strings.Replace(templateTpl, "{{tpl - $tags}}", tag, 1)
			item = strings.ReplaceAll(item, "{{tpl - $title}}", f.Title)
			item = strings.ReplaceAll(item, "{{tpl - $name}}", f.Name)
			columns.WriteString(item)
			hasDict = true
		} else if f.Type == "复选框" || f.Type == "下拉框多选" {
			content = strings.Replace(content, "<Tpl - Import - Dict>", "import { Dict, reqDict } from '@/api/basic/dict';", 1)
			content = strings.Replace(content, "<Tpl - Import - Dict - Init>", "requestDict();", 1)
			content = strings.Replace(content, "<Tpl - Request - Dict - Object>", dictRequest, 1)

			kks := fmt.Sprintf("dictMap['%s'][item]", strings.TrimSpace(f.DictType))
			tag := fmt.Sprintf("<a-space wrap><a-tag v-for=\"item in record.%s\" :color=\"%s.color\">{{%s.label}}</a-tag></a-space>", f.Name, kks, kks)
			item := strings.Replace(templateTpl, "{{tpl - $tags}}", tag, 1)
			item = strings.ReplaceAll(item, "{{tpl - $title}}", f.Title)
			item = strings.ReplaceAll(item, "{{tpl - $name}}", f.Name)
			columns.WriteString(item)
			hasDict = true
		} else if f.Type == "开关" {
			columns.WriteString(strings.Replace(strings.Replace(switchTpl, "{{tpl - $title}}", f.Title, 1), "{{tpl - $name}}", f.Name, 1))
		} else if strings.Contains(f.Type, "上传图片") {
			imageTag := fmt.Sprintf("<a-space wrap><a-image v-for=\"url in record.%s\" width=\"50\" :src=\"String(url).includes('http') ? url : urlPrefix + 'admin/' + url\"></a-image></a-space>", f.Name)
			item := strings.Replace(templateTpl, "{{tpl - $tags}}", imageTag, 1)
			item = strings.ReplaceAll(item, "{{tpl - $title}}", f.Title)
			item = strings.ReplaceAll(item, "{{tpl - $name}}", f.Name)
			columns.WriteString(item)
		} else if f.Type == "上传视频" || f.Type == "上传文件" {
			tag := fmt.Sprintf("<a-space wrap><a-link v-for=\"(url, index) in record.%s\" :href=\"String(url).includes('http') ? url : urlPrefix + 'admin/' + url\">链接 - {{ index }}</a-link></a-space>", f.Name)

			item := strings.Replace(templateTpl, "{{tpl - $tags}}", tag, 1)
			item = strings.ReplaceAll(item, "{{tpl - $title}}", f.Title)
			item = strings.ReplaceAll(item, "{{tpl - $name}}", f.Name)
			columns.WriteString(item)
		} else if strings.Contains(f.Type, "富文本") {
			columns.WriteString("")
		} else if strings.Contains(f.Type, "text") {
			columns.WriteString(six.Strings(`                    <a-table-column title="`, f.Title, `" data-index="`, f.Name, `"><template #cell="{ record }"><a-typography-text ellipsis>{{ record.`, f.Name, ` }}</a-typography-text></template></a-table-column>`).String())
		} else {
			columns.WriteString(six.Strings(`                    <a-table-column title="`, f.Title, `" data-index="`, f.Name, `"></a-table-column>`, "\n").String())
		}
	}
	content = strings.ReplaceAll(content, "<Tpl-Table-Columns>", strings.TrimSpace(columns.String()))
	content = strings.Replace(content, "<Tpl - Dict>", utils.Ternary(hasDict, `:dict="dictList"`, ""), 1)
	content = strings.Replace(content, "<Tpl - Request - Dict - Init>", utils.Ternary(hasDict, `requestDict()`, ""), 1)

	path := six.Str(ProjectFrontRootPath).Append("/src/views/codegen/", six.Str(c.v.Table).FirstUpper().RemoveSpace()).String()
	componentsPath := path + "/components"

	if err = os.MkdirAll(componentsPath, os.ModePerm); err != nil {
		c.err = err
	}

	c.err = os.WriteFile(path+"/index.vue", []byte(content), 0666)
	return c
}

func (c *CodeGenerator) form() *CodeGenerator {
	if c.err != nil {
		return c
	}
	if c.err != nil {
		return c
	}
	tplFilepath := six.Str(ProjectServerRootPath).Append("/static/generator_tpl/front/form.vue.tpl").String()
	file, err := os.ReadFile(tplFilepath)
	if err != nil {
		c.err = err
		return c
	}

	content := string(file)
	content = strings.ReplaceAll(content, "<Tpl-Module>", six.Str(c.v.Table).Snake2BigCamel())
	content = strings.ReplaceAll(content, "<Tpl-Module-Lower>", six.Str(c.v.Table).RemoveSpace())

	//var fields strings.Builder
	itemTpl := `            
        <a-form-item field="{{tpl: $field}}" label="{{tpl: $label}}"{{tpl: $required}}>
            {{tpl - $input}}
        </a-form-item>`

	inputTpl := `<a-input v-model.trim="form.{{tpl - $name}}" placeholder="请输入{{tpl: $label}}"></a-input>`
	inputNumberTpl := `<a-input-number v-model.trim="form.{{tpl - $name}}" placeholder="请输入{{tpl: $label}}"></a-input-number>`

	radioTpl := `
<a-radio-group v-model="form.{{tpl - $name}}" type="button" :size="size">
	<a-radio v-for="item in dict['{{tpl - $dict}}']" :value="item.value">{{ item.label }}</a-radio>
</a-radio-group>`

	selectTpl := `<a-select v-model="form.{{tpl - $name}}" :options="dict['{{tpl - $dict}}']" placeholder="请选择{{tpl - $label}}" allow-clear allow-search {{tpl - $multiple}}/>`

	checkboxTpl := `<a-checkbox-group v-model="form.{{tpl - $name}}" :options="dict['{{tpl - $dict}}']" />`

	switchTpl := `<a-switch type="round" v-model="form.{{tpl - $name}}" checked-text=" 正 常 " unchecked-text=" 禁 用 "></a-switch>`

	imageTpl := `
	<a-upload list-type="picture-card" multiple
		accept="image/*"
		:action="uploadUrl"
		:headers="uploadHeaders"
		@success="uploadSuccess{{tpl - $name}}"
		@error="uploadError{{tpl - $name}}"
		@before-remove="beforeRemove{{tpl - $name}}"
		v-model="uploads_{{tpl - $name}}"
		image-preview
	/>`

	imageOneTpl := `
	<a-upload
		accept="image/*"
		:show-file-list="false"
		:action="uploadUrl"
		:headers="uploadHeaders"
		@success="uploadSuccess{{tpl - $name}}"
		@error="uploadError{{tpl - $name}}"
	/>
	<a-input v-model="form.{{tpl - $name}}" readonly></a-input>
`

	fileTpl := `
	<a-upload multiple
		{{tpl - $video}}
		:action="uploadUrl"
		:headers="uploadHeaders"
		v-model="uploads_{{tpl - $name}}"
		@success="uploadSuccess{{tpl - $name}}"
		@error="uploadError{{tpl - $name}}"
		@before-remove="beforeRemove{{tpl - $name}}"
	>
		<template #upload-button>
			<a-button type="primary" :size="size">
				<template #icon>
					<icon-upload />
				</template>
				<template #default>上传文件</template>
			</a-button>
		</template>
	</a-upload>
`

	uploadJs := `
	const uploads_{{tpl - $name}} = ref([])
	const uploadSuccess{{tpl - $name}} = async (e) => {
		form.value.{{tpl - $name}}.push(e.response.data.url)
		uploads_{{tpl - $name}}.value.push(e)
    }

	const uploadError{{tpl - $name}} = async (e) => {
        Message.error('上传失败');
    }
	
	const beforeRemove{{tpl - $name}} = async (e) => {
		return new Promise((resolve, reject) => {
			Modal.confirm({
				title: 'on-before-remove',
				content: "确认删除？",
				onOk: () => {
					form.value.{{tpl - $name}} = form.value.{{tpl - $name}}.filter(item => item != e.response.data.url)
					uploads_{{tpl - $name}}.value = uploads_{{tpl - $name}}.value.filter(item => item != e)
                    resolve(true)
				},
				onCancel: () => reject('cancel'),
			});
		});
    }
`
	uploadOneJs := `
	const upload_{{tpl - $name}} = ref([])
	const uploadSuccesste_image_one = async (e) => {
		form.value.te_image_one= [e.response.data.url]
		upload_{{tpl - $name}}.value= [e]
	};

	const uploadError{{tpl - $name}} = async (e) => {
        Message.error('上传失败');
    };
`

	var hasDict bool
	var hasEditor bool
	var items []string
	var jsArr []string
	for _, f := range *c.v.Fields {
		var input string
		if f.Type == "单选框" {
			hasDict = true
			input = strings.ReplaceAll(radioTpl, "{{tpl - $name}}", f.Name)
			input = strings.ReplaceAll(input, "{{tpl - $dict}}", strings.TrimSpace(f.DictType))
		} else if f.Type == "复选框" {
			hasDict = true
			input = strings.ReplaceAll(checkboxTpl, "{{tpl - $name}}", f.Name)
			input = strings.ReplaceAll(input, "{{tpl - $dict}}", strings.TrimSpace(f.DictType))
		} else if f.Type == "下拉框" {
			input = strings.ReplaceAll(selectTpl, "{{tpl - $name}}", f.Name)
			input = strings.ReplaceAll(input, "{{tpl - $label}}", f.Title)
			input = strings.ReplaceAll(input, "{{tpl - $dict}}", strings.TrimSpace(f.DictType))
			input = strings.ReplaceAll(input, "{{tpl - $multiple}}", "")
		} else if f.Type == "下拉框多选" {
			hasDict = true
			input = strings.ReplaceAll(selectTpl, "{{tpl - $name}}", f.Name)
			input = strings.ReplaceAll(input, "{{tpl - $label}}", f.Title)
			input = strings.ReplaceAll(input, "{{tpl - $dict}}", strings.TrimSpace(f.DictType))
			input = strings.ReplaceAll(input, "{{tpl - $multiple}}", "multiple")
		} else if f.Type == "开关" {
			input = strings.ReplaceAll(switchTpl, "{{tpl - $name}}", f.Name)
		} else if f.Type == "上传图片" {
			input = strings.ReplaceAll(imageOneTpl, "{{tpl - $name}}", f.Name)
			jsArr = append(jsArr, strings.ReplaceAll(uploadOneJs, "{{tpl - $name}}", f.Name))
		} else if f.Type == "上传图片（多图）" {
			input = strings.ReplaceAll(imageTpl, "{{tpl - $name}}", f.Name)
			jsArr = append(jsArr, strings.ReplaceAll(uploadJs, "{{tpl - $name}}", f.Name))
		} else if f.Type == "上传文件" || f.Type == "上传视频" {
			input = strings.ReplaceAll(fileTpl, "{{tpl - $name}}", f.Name)
			input = strings.ReplaceAll(input, "{{tpl - $video}}", utils.Ternary(f.Type == "上传视频", `accept="video/*"`, ""))
			jsArr = append(jsArr, strings.ReplaceAll(uploadJs, "{{tpl - $name}}", f.Name))
		} else if f.Type == "时间选择器" {
			input = fmt.Sprintf("<a-time-picker v-model='form.%s' />", f.Name)
		} else if f.Type == "日期选择器" {
			input = fmt.Sprintf("<a-date-picker v-model='form.%s'/>", f.Name)
		} else if f.Type == "日期时间选择器" {
			input = fmt.Sprintf("<a-date-picker v-model='form.%s' format='YYYY-MM-DD HH:mm:ss' show-time/>", f.Name)
		} else if f.Type == "富文本" {
			hasEditor = true
			input = fmt.Sprintf("<wang @change='(e) => form.%s = e'></wang>", f.Name)
		} else if f.Type == "text" {
			input = fmt.Sprintf("<a-textarea v-model='form.%s' show-word-limit auto-size placeholder='请输入 ...' allow-clear/>", f.Name)
		} else if strings.Contains(f.Type, "int") || strings.Contains(f.Type, "float") || strings.Contains(f.Type, "decimal") {
			input = strings.ReplaceAll(inputNumberTpl, "{{tpl - $name}}", f.Name)
			input = strings.ReplaceAll(input, "{{tpl: $label}}", f.Title)
		} else {
			input = strings.ReplaceAll(inputTpl, "{{tpl - $name}}", f.Name)
			input = strings.ReplaceAll(input, "{{tpl: $label}}", f.Title)
		}

		items = append(items, strings.Replace(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(itemTpl, "{{tpl: $field}}", f.Name), "{{tpl: $label}}", f.Title), "{{tpl: $required}}", utils.Ternary(f.IsRequired, " required", "")), "{{tpl - $input}}", input, 1))
	}

	if hasEditor {
		content = strings.Replace(content, "<Tpl - Editor - Import>", "import wang from '@/components/editor/wang.vue'", 1)
	} else {
		content = strings.Replace(content, "<Tpl - Editor - Import>", "", 1)
	}

	content = strings.ReplaceAll(content, "<Tpl-From-Items>", strings.TrimSpace(strings.Join(items, "\n")))
	content = strings.Replace(content, "<Tpl - ImageList - Success - Error>", strings.Join(jsArr, "\n\n"), 1)
	content = strings.Replace(content, "<Tpl - Dict>", utils.Ternary(hasDict, "dict:Object;", ""), 1)

	path := six.Str(ProjectFrontRootPath).Append("/src/views/codegen/", six.Str(c.v.Table).RemoveSpace(), "/components/form-save.vue").String()
	c.err = os.WriteFile(path, []byte(content), 0666)
	return c
}

func (c *CodeGenerator) finished() error {
	if c.err != nil {
		return c.err
	}

	var bros []models.AuthRule
	if err := db.DB().Where("parent_id = ?", c.v.ParentModule).Find(&bros).Error; err != nil {
		return err
	}

	var broLastId int64
	var broLastOrder int
	if len(bros) > 0 {
		for _, b := range bros {
			if b.Id > broLastId {
				broLastId = b.Id
			}

			if b.Order > broLastOrder {
				broLastOrder = b.Order
			}
		}
	}

	if broLastId == 0 {
		broLastId = six.Str(fmt.Sprintf("%d01", c.v.ParentModule)).Int64()
	}

	if err := db.DB().Where("id in ?", []int64{broLastId,
		six.Str(fmt.Sprintf("%d01", broLastId)).Int64(),
		six.Str(fmt.Sprintf("%d02", broLastId)).Int64(),
		six.Str(fmt.Sprintf("%d03", broLastId)).Int64(),
		six.Str(fmt.Sprintf("%d04", broLastId)).Int64(),
		six.Str(fmt.Sprintf("%d05", broLastId)).Int64(),
		six.Str(fmt.Sprintf("%d06", broLastId)).Int64(),
		six.Str(fmt.Sprintf("%d07", broLastId)).Int64(),
	}).Unscoped().Delete(&models.AuthRule{}).Error; err != nil {
		return err
	}

	var rules []models.AuthRule
	rg := six.Str(strings.TrimSpace(c.v.Table)).Lower().String()
	rules = append(rules, models.AuthRule{
		MODEL: db.MODEL{
			Id: broLastId,
		},
		SoftDelete: db.SoftDelete{},
		ParentId:   c.v.ParentModule,
		Type:       models.RuleTypesMenu,
		Component:  "",
		Name:       six.Str(strings.TrimSpace(c.v.Table)).Snake2BigCamel(),
		Path:       six.Str(strings.TrimSpace(c.v.Table)).Snake2LittleCamel(),
		Auth:       c.auth,
		Locale:     c.v.Title,
		Icon:       "",
		Order:      broLastOrder + 1,
		Status:     1,
	})

	rules = append(rules, models.AuthRule{
		MODEL: db.MODEL{
			Id: six.Str(fmt.Sprintf("%d01", broLastId)).Int64(),
		},
		SoftDelete: db.SoftDelete{},
		ParentId:   broLastId,
		Type:       models.RuleTypesPage,
		Component:  "",
		Auth:       c.auth + ":list",
		ApiRoute:   rg + "/list",
		Locale:     "列表（分页）",
		Icon:       "",
		Order:      1,
		Status:     1,
	})

	if c.IsTree {
		rules = append(rules, models.AuthRule{
			MODEL: db.MODEL{
				Id: six.Str(fmt.Sprintf("%d02", broLastId)).Int64(),
			},
			SoftDelete: db.SoftDelete{},
			ParentId:   broLastId,
			Type:       models.RuleTypesPage,
			Component:  "",
			Auth:       c.auth + ":tree-select",
			ApiRoute:   rg + "/tree-select",
			Locale:     "列表（树）",
			Icon:       "",
			Order:      2,
			Status:     1,
		})
	}

	rules = append(rules, models.AuthRule{
		MODEL: db.MODEL{
			Id: six.Str(fmt.Sprintf("%d03", broLastId)).Int64(),
		},
		SoftDelete: db.SoftDelete{},
		ParentId:   broLastId,
		Type:       models.RuleTypesPage,
		Component:  "",
		Auth:       c.auth + ":select",
		ApiRoute:   rg + "/select",
		Locale:     "列表（选择）",
		Icon:       "",
		Order:      3,
		Status:     1,
	})

	rules = append(rules, models.AuthRule{
		MODEL: db.MODEL{
			Id: six.Str(fmt.Sprintf("%d04", broLastId)).Int64(),
		},
		SoftDelete: db.SoftDelete{},
		ParentId:   broLastId,
		Type:       models.RuleTypesButton,
		Component:  "",
		Auth:       c.auth + ":get",
		ApiRoute:   rg + "/get",
		Locale:     "详情",
		Icon:       "",
		Order:      4,
		Status:     1,
	})

	rules = append(rules, models.AuthRule{
		MODEL: db.MODEL{
			Id: six.Str(fmt.Sprintf("%d05", broLastId)).Int64(),
		},
		SoftDelete: db.SoftDelete{},
		ParentId:   broLastId,
		Type:       models.RuleTypesButton,
		Component:  "",
		Auth:       c.auth + ":add",
		ApiRoute:   rg + "/add",
		Locale:     "添加",
		Icon:       "",
		Order:      5,
		Status:     1,
	})

	rules = append(rules, models.AuthRule{
		MODEL: db.MODEL{
			Id: six.Str(fmt.Sprintf("%d06", broLastId)).Int64(),
		},
		SoftDelete: db.SoftDelete{},
		ParentId:   broLastId,
		Type:       models.RuleTypesButton,
		Component:  "",
		Auth:       c.auth + ":save",
		ApiRoute:   rg + "/save",
		Locale:     "编辑",
		Icon:       "",
		Order:      6,
		Status:     1,
	})

	rules = append(rules, models.AuthRule{
		MODEL: db.MODEL{
			Id: six.Str(fmt.Sprintf("%d07", broLastId)).Int64(),
		},
		SoftDelete: db.SoftDelete{},
		ParentId:   broLastId,
		Type:       models.RuleTypesButton,
		Component:  "",
		Auth:       c.auth + ":del",
		ApiRoute:   rg + "/del",
		Locale:     "删除",
		Icon:       "",
		Order:      7,
		Status:     1,
	})
	return db.DB().Create(&rules).Error
}
