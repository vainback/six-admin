package authService

import (
	"fmt"
	"github.com/vainback/six-util/v3"
	"golang.org/x/exp/maps"
	"log"
	"six-go/database/db"
	"six-go/database/models"
	"six-go/extra/xcache"
	"slices"
	"strings"
)

func HasAuth(route string, roleId ...int64) bool {
	rightRoleIds := getRightRoleIdsWithRoleIds(roleId...)
	if len(rightRoleIds) == 0 {
		return false
	}
	ruleIds := GetRuleIdsByRoles(rightRoleIds...)
	if len(ruleIds) == 0 {
		return false
	}

	rules := GetRules()
	if len(rules) == 0 {
		return false
	}

	for _, rule := range rules {
		if rule.ApiRoute == route && rule.Status == 1 && slices.Contains(ruleIds, rule.Id) {
			return true
		}
	}
	return false
}

func GetRuleIdsByRoles(roleId ...int64) []int64 {
	if len(roleId) == 0 {
		return nil
	}
	var ids []int64
	for _, v := range roleId {
		ids = append(ids, getRuleIdsByRole(v)...)
	}
	return ids
}

func GetNamesByRoute(route string) string {
	rules := GetRules()
	if len(rules) == 0 {
		return ""
	}
	var rule models.AuthRule
	for _, v := range rules {
		if v.ApiRoute == route {
			rule = v
			break
		}
	}

	if rule.Auth == "" {
		return ""
	}

	var auths []string
	for index, val := range strings.Split(rule.Auth, ":") {
		if index == 0 {
			auths = append(auths, val)
		} else {
			//auths = append(auths, utils.Strings(auths[index-1], ":", val).String())
			auths = append(auths, six.Str(auths[index-1]).Append(":", val).String())
		}
	}
	var names = make(map[int64]string)
	for _, val := range rules {
		if slices.Contains(auths, val.Auth) {
			names[val.Id] = val.Locale
		}
	}

	ids := maps.Keys(names)
	slices.Sort(ids)

	var nameArr = make([]string, len(ids))
	for i, id := range ids {
		nameArr[i] = names[id]
	}
	return strings.Join(nameArr, "--")
}

func getRuleIdsByRole(roleId int64) []int64 {
	key := fmt.Sprintf("auth_relation_%d", roleId)
	ruleIds := xcache.SyncLoad[[]int64](key)
	if len(ruleIds) == 0 {
		if err := db.DB().Table(models.TableAuthRelation.TableName()).Select("rule_id").Where("role_id = ?", roleId).Scan(&ruleIds).Error; err != nil {
			return nil
		}
		xcache.SyncStore(key, ruleIds)
	}
	return ruleIds
}

func GetRules() []models.AuthRule {
	key := "auth_rule"
	rules := xcache.SyncLoad[[]models.AuthRule](key)
	if len(rules) == 0 {
		if err := db.DB().Find(&rules).Error; err != nil {
			return nil
		}
		xcache.SyncStore(key, rules)
	}
	return rules
}

func GetRoles() []models.AuthRole {
	key := "auth_role"
	roles := xcache.SyncLoad[[]models.AuthRole](key)
	if len(roles) == 0 {
		if err := db.DB().Find(&roles).Error; err != nil {
			return nil
		}
		xcache.SyncStore(key, roles)
	}
	return roles
}

func getRightRoleIdsWithRoleIds(ids ...int64) []int64 {
	if len(ids) == 0 {
		return nil
	}

	roles := GetRoles()
	if len(roles) == 0 {
		return nil
	}
	var rightIds []int64
	for _, v := range roles {
		if v.Status == 1 && slices.Contains(ids, v.Id) {
			rightIds = append(rightIds, v.Id)
		}
	}
	return rightIds
}

func LoadAllRelation() {
	var list []models.AuthRelation
	if err := db.DB().Find(&list).Error; err != nil {
		log.Println(err)
		return
	}

	if len(list) == 0 {
		return
	}

	var relations = make(map[int64][]int64)
	for _, v := range list {
		if _, ok := relations[v.RoleId]; !ok {
			relations[v.RoleId] = make([]int64, 0)
		}
		relations[v.RoleId] = append(relations[v.RoleId], v.RuleId)
	}

	for role, rules := range relations {
		xcache.SyncStore(fmt.Sprintf("auth_relation_%d", role), rules)
	}
}
