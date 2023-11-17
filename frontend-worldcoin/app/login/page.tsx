"use client"
import { useState } from "react";
import { CredentialType, IDKitWidget, ISuccessResult,useIDKit } from '@worldcoin/idkit'
import { useRouter } from "next/navigation";

declare global {
    interface Window {
        ethereum?: any;
    }
}

export const formatBalance = (rawBalance: string) => {
    const balance = (parseInt(rawBalance) / 1000000000000000000).toFixed(2)
    return balance
}

export const formatChainAsNum = (chainIdHex: string) => {
    const chainIdNum = parseInt(chainIdHex)
    return chainIdNum
}

const Login = () => {
    const [wordCoinAddress, setWordCoinAddress] = useState("")
    const { open, setOpen } = useIDKit()

    const router = useRouter()

    const onSuccess = (data: ISuccessResult) => {
        console.log(data);
        setWordCoinAddress(data.nullifier_hash)
    }

    return (
        <div className="bg-white py-6 sm:py-8 lg:py-12">
            <div className="mx-auto max-w-lg">
                <IDKitWidget
                    app_id={process.env.NEXT_PUBLIC_WORLD_COIN_ID}// obtained from the Developer Portal
                    action="login_eth" // this is your action name from the Developer Portal
                    onSuccess={onSuccess} // callback when the modal is closed
                    credential_types={[CredentialType.Orb, CredentialType.Phone]} // optional, defaults to ['orb']
                >
                    {({ open }) =>
                        <div className="flex flex-col gap-4 p-4 md:p-8">
                            <button onClick={open} className="block rounded-lg bg-gradient-to-r from-green-400 via-green-500 to-green-600 px-8 py-3 text-center text-sm font-semibold text-white outline-none ring-gray-300 transition duration-100 hover:bg-gray-700 focus-visible:ring active:bg-gray-600 md:text-base">Verify with World ID</button>
                        </div>
                    }
                </IDKitWidget>

            </div>
        </div>
    )
}

export default Login;