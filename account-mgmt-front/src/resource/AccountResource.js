import axios from 'axios';

const API_BASE_PATH = 'http://localhost:3000/api/account-mgmt';

const GetAccountBalance = async () => {
  let result = { balance: 0 };
  try {
    const uri = `${API_BASE_PATH}/account/balance`;
    const resp = await axios.get(uri);
    result = resp.data;
  } catch(e) {
    console.log(e);
    return result
  }
  return result
}

export default GetAccountBalance;