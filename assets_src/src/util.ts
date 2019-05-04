export function getCookieByName(name: string): string {
  const { cookie } = document;
  // console.log(cookie);
  const reg = new RegExp(`${name}=([^;]*)`);
  const val = cookie.replace(reg, '$1');
  return val;
}
