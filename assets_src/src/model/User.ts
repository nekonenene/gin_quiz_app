interface User {
  id: number;
  name: string;
  email: string;
}

export const defaultUser: User = {
  id: 0,
  name: '',
  email: '',
}

export default User;
