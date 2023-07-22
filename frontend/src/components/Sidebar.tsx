import React from "react";
import { ethers } from "ethers";
import useIsMetaMaskInstalled from "../useIsMetaMaskInstalled";

interface Props {
  setAccount: React.Dispatch<React.SetStateAction<string | undefined>>;
  account?: string;
  contractAddress: string;
}

const Sidebar = ({ setAccount, account, contractAddress }: Props) => {
  const isMetaMaskInstalled = useIsMetaMaskInstalled();

  const handleOnConnect = () => {
    window.ethereum
      .request({ method: "eth_requestAccounts" })
      .then((accounts: string[]) => {
        setAccount(ethers.utils.getAddress(accounts[0]));
      })
      .catch((err: any) => console.log(err));
  };

  return (
    <div className="sidebar">
      {account && (
        <>
          <b>Connected as:</b>
          <br />
          <small><a target="blank" href={"http://etherscan.io/address/" + account}> {account}</a></small>
        </>
      )}
      {!account && (
        <button onClick={handleOnConnect} disabled={!isMetaMaskInstalled}>
          Connect your wallet
        </button>
      )}
      {!isMetaMaskInstalled && <p>Please install MetaMask</p>}
      <h1>Wedding DApp</h1>
      <p>
      <svg aria-hidden="true" height="16" viewBox="0 0 16 16" version="1.1" width="16" data-view-component="true" className="link">
          <path d="m7.775 3.275 1.25-1.25a3.5 3.5 0 1 1 4.95 4.95l-2.5 2.5a3.5 3.5 0 0 1-4.95 0 .751.751 0 0 1 .018-1.042.751.751 0 0 1 1.042-.018 1.998 1.998 0 0 0 2.83 0l2.5-2.5a2.002 2.002 0 0 0-2.83-2.83l-1.25 1.25a.751.751 0 0 1-1.042-.018.751.751 0 0 1-.018-1.042Zm-4.69 9.64a1.998 1.998 0 0 0 2.83 0l1.25-1.25a.751.751 0 0 1 1.042.018.751.751 0 0 1 .018 1.042l-1.25 1.25a3.5 3.5 0 1 1-4.95-4.95l2.5-2.5a3.5 3.5 0 0 1 4.95 0 .751.751 0 0 1-.018 1.042.751.751 0 0 1-1.042.018 1.998 1.998 0 0 0-2.83 0l-2.5 2.5a1.998 1.998 0 0 0 0 2.83Z"></path>
      </svg>
        <a target="blank" href={"http://etherscan.io/address/" + contractAddress}> View contract on Etherscan</a>
      </p>
    </div>
  );
};

export default Sidebar;
