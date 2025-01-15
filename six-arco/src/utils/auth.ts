const TOKEN_KEY = 'token';

const isLogin = () => {
  return !!localStorage.getItem(TOKEN_KEY);
};

const getToken = () => {
  return localStorage.getItem(TOKEN_KEY);
};

const setToken = (token: string) => {
  localStorage.setItem(TOKEN_KEY, token);
};

const clearToken = () => {
  localStorage.removeItem(TOKEN_KEY);
};

const getTenantId = () => {
  return localStorage.getItem("login.tenant_id");
}

const isTenantRoot = () => {
  return Number(getTenantId()) === 1;
}

export { isLogin, getToken, setToken, clearToken, getTenantId, isTenantRoot };
