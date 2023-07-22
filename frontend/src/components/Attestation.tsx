import React, { useEffect, useState } from 'react';
import { EAS, SchemaEncoder } from "@ethereum-attestation-service/eas-sdk";
import { ethers } from 'ethers';
import Wedding from "../contract/contract.json";

export default function Attestation() {
    const [newAttestationUID, setNewAttestationUID] = useState('');

    const provider = new ethers.providers.Web3Provider(window.ethereum);
    const signer = provider.getSigner();
    useEffect(() => {
        const EASContractAddress = "0xC2679fBD37d54388Ce493F1DB75320D236e1815e"; // Sepolia v0.26
        const weddingContractAddress = '0x3a405b3f52de981b85c2e2db5e5bacd503ea6538'; // Your wedding contract address
        const abi = Wedding.abi;
        const weddingContract = new ethers.Contract(weddingContractAddress, abi, provider);
        const eas = new EAS(EASContractAddress);
        
        // then pass to eas.connect()
        eas.connect(signer);

        const schemaEncoder = new SchemaEncoder("address Address1, address Address2, string Name1, string Name2");

        const onNewSignature = async () => {
            console.log('NewSignature event fired!');
        };
        weddingContract.on('NewSignature', onNewSignature);

        const onMarriageDone = async () => {
            console.log('MarriageDone event fired!');
            const [fiance1WalletAddr, fiance1Name] = await weddingContract.getFiance1();
            const [fiance2WalletAddr, fiance2Name] = await weddingContract.getFiance2();

            const encodedData = schemaEncoder.encodeData([
                { name: "Address1", value: fiance1WalletAddr, type: "address" },
                { name: "Address2", value: fiance2WalletAddr, type: "address" },
                { name: "Name1", value: fiance1Name, type: "string" },
                { name: "Name2", value: fiance2Name, type: "string" },
            ]);

            const schemaUID = "0x7bcf2d70b6424e5038dbdd98f7c37a8debb96ebfc2b711a2c4a5272ef0c201b2";

            const tx = await eas.attest({
                schema: schemaUID,
                data: {
                    recipient: "0x895c6b812E276FfBba21fE560327b4C5Aa1c17f3",
                    expirationTime: 0,
                    revocable: true,
                    data: encodedData,
                },
            });

            const newAttestationUID = await tx.wait();
            setNewAttestationUID(newAttestationUID);
        };
        weddingContract.on('MarriageDone', onMarriageDone);

        return () => {
            // Clean up listeners when the component unmounts
            weddingContract.off('NewSignature', onNewSignature);
            weddingContract.off('MarriageDone', onMarriageDone);
        };
    }, [provider, signer]);  // dependencies for useEffect

    return (
        <div>
            New attestation UID: {newAttestationUID}
        </div>
    );
}
