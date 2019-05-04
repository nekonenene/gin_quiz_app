export function getCookieByName(name: string): string {
  var cookie = document.cookie;
  // console.log(cookie);
  var reg = new RegExp(`${name}=([^;]*)`);
  var val = cookie.replace(reg, '$1');
  return val;
}
