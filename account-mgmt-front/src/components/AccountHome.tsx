import React, { useState, useEffect } from "react";
import Accordion from "react-bootstrap/Accordion";
import Button from "react-bootstrap/Button";
import Card from "react-bootstrap/Card";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faUser,
  faArrowUp,
  faArrowDown,
  faDollarSign,
} from "@fortawesome/free-solid-svg-icons";
import "bootswatch/dist/lumen/bootstrap.min.css";
import GetTransactions from "../resource/TransactionResource";
import GetAccountBalance from "../resource/AccountResource";

interface Account {
  balance: number;
  name: string;
}
interface AccountBalance {
  balance: number;
}

interface Transaction {
  type: string;
  amount: number;
  id: string;
  effectiveDate: string;
}

const isCredit = (tx: Transaction) => tx.type === "credit";

const getColor = (tx: Transaction) => (isCredit(tx) ? "green" : "red");

function AccountHome(props: any): JSX.Element {
  const [account, setAccount] = useState<Account>({
    balance: -1,
    name: "L. Dzisiuk",
  });
  const [transactions, setTransactions] = useState<Transaction[]>([]);

  useEffect(() => {
    GetTransactions().then((transactions: Transaction[]) =>
      setTransactions(transactions)
    );
  }, [account]);

  useEffect(() => {
    GetAccountBalance().then((ab: AccountBalance) =>
      setAccount({ ...account, ...ab })
    );
  }, []);

  return (
    <div className="container p-4">
      <div className="row">
        <div className="col-md-6 offset-md-3">
          <div className="card">
            <div className="card-body">
              <div className="row">
                <h1 className="m-2 ml-3">
                  <FontAwesomeIcon icon={faUser} className="m-1" />
                  {account.name}
                </h1>
                <h3 className="m-4 pull-right">
                  Balance:{" "}
                  {account.balance < 0 ? "...loading" : account.balance}
                </h3>
              </div>
            </div>
          </div>
          {transactions.map((tx: Transaction, i: number) => {
            return (
              <Accordion>
                <Card>
                  <Card.Header>
                    <Accordion.Toggle
                      variant="link"
                      as={Button}
                      eventKey="0"
                    >
                      <h3 className="row m-3" key={i}>
                        {tx.amount}
                        <FontAwesomeIcon
                          color={getColor(tx)}
                          icon={isCredit(tx) ? faArrowUp : faArrowDown}
                          className="m-1"
                        />
                        <FontAwesomeIcon
                          color={getColor(tx)}
                          icon={faDollarSign}
                          className="m-1"
                        />
                      </h3>
                    </Accordion.Toggle>
                  </Card.Header>
                  <Accordion.Collapse eventKey="0">
                    <Card.Body><span>
                      <p>Transaction ID: {tx.id}</p>
                      <p>Date: {tx.effectiveDate}</p>
                      </span></Card.Body>
                  </Accordion.Collapse>
                </Card>
              </Accordion>
            );
          })}
        </div>
      </div>
    </div>
  );
}

export default AccountHome;
