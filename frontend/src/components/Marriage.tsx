// Marriage.tsx
import React, { useState } from 'react';
import { ethers } from 'ethers';
import Wedding from "../contract/contract.json";

export default function Marriage() {
    const [wallet, setWallet] = useState<ethers.providers.JsonRpcSigner | null>(null);
    const provider = new ethers.providers.Web3Provider(window.ethereum);

    const weddingContractAddress = '0x4fcd2d1d4175262f110b0c7a055ea25c2aed656c'; // Your wedding contract address
    const abi = Wedding.abi;
    const contract = new ethers.Contract(weddingContractAddress, abi, provider);

    // Connect wallet button was clicked
    const connectWallet = async () => {
        const accounts = await window.ethereum.request({ method: 'eth_requestAccounts' });
        const signer = provider.getSigner(accounts[0]);
        setWallet(signer);
    }

    const sayYes = async () => {
        if (wallet === null) {
            window.alert("Please connect your wallet.");
            return;
        }

        // Assure TypeScript that wallet isn't null by adding a null check
        if (wallet) {
            const contractWithSigner = contract.connect(wallet);
            const tx = await contractWithSigner.sayYes();
            await tx.wait(); // wait for the transaction to be mined

            window.alert('You said yes!');
        }
    }

    return (
        <div>
            <button onClick={connectWallet}>Connect Wallet</button>
            <button onClick={sayYes}>Say Yes</button>
        </div>
    );
}
