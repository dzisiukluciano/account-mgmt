import axios from 'axios';

const API_BASE_PATH = 'http://localhost:3000/api/account-mgmt';

const GetTransactions = async () => {
  let transactions = [];
  try {
    const uri = `${API_BASE_PATH}/account/transactions`;
    const resp = await axios.get(uri);
    transactions = resp.data;
  } catch(e) {
    console.log(e);
    return transactions
  }
  return transactions
}

export default GetTransactions;