import React, { useEffect, useState } from 'react';
import { EAS, SchemaEncoder } from "@ethereum-attestation-service/eas-sdk";
import { ethers } from 'ethers';
import Wedding from "../contract/contract.json";

export default function Attestation() {
    const [newAttestationUID, setNewAttestationUID] = useState('');

    useEffect(() => {
        (async function attest() {
            const provider = new ethers.providers.JsonRpcProvider("https://rpc.sepolia.dev");

            const EASContractAddress = "0xC2679fBD37d54388Ce493F1DB75320D236e1815e"; // Sepolia v0.26
            const weddingContractAddress = '0x8DD43110D79D51d29155a94e2f73754F86f13981'; // Your wedding contract address
            const abi = Wedding.abi;
            const weddingContract = new ethers.Contract(weddingContractAddress, abi, provider);
            const eas = new EAS(EASContractAddress);

            const getLastSigner = async () => {
                const lastSignerAddress = await weddingContract.getLastSigner();
                return lastSignerAddress;
            }

            eas.connect(await getLastSigner());

            // Initialize SchemaEncoder with the schema string
            const schemaEncoder = new SchemaEncoder("address Address1, address Address2, string Name1, string Name2");

            // Fetch fiance details:
            const getFianceDetails = async () => {
                const [fiance1WalletAddr, fiance1Name] = await weddingContract.getFiance1();
                const [fiance2WalletAddr, fiance2Name] = await weddingContract.getFiance2();
                return {
                    Address1: fiance1WalletAddr,
                    Name1: fiance1Name,
                    Address2: fiance2WalletAddr,
                    Name2: fiance2Name
                };
            }

            const {Address1, Name1, Address2, Name2} = await getFianceDetails();

            const encodedData = schemaEncoder.encodeData([
                { name: "Address1", value: Address1, type: "address" },
                { name: "Address2", value: Address2, type: "address" },
                { name: "Name1", value: Name1, type: "string" },
                { name: "Name2", value: Name2, type: "string" },
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
        })();
    }, []);

    return (
        <div>
            New attestation UID: {newAttestationUID}
        </div>
    );
}
