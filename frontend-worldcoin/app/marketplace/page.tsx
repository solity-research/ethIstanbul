import Link from "next/link";

const Marketplace = () => {
    return (
        <>
            <div className="mb-6 flex flex-col gap-4 sm:mb-8 md:gap-6">
                <div className="w-full flex justify-end">
                    <Link href="/marketplace/add" className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded inline-flex items-center">
                        <span>+ Add your Protocol GPT</span>
                    </Link>
                </div>
                <div className="flex flex-wrap gap-x-4 overflow-hidden rounded-lg border sm:gap-y-4 lg:gap-6">
                    <a href="#" className="group relative block h-48 w-32 overflow-hidden bg-gray-100 sm:h-56 sm:w-60">
                        <img src="./aave_log.png" loading="lazy" alt="Photo by Thái An" className="h-full w-full object-cover object-center transition duration-200 group-hover:scale-110" />
                    </a>

                    <div className="flex flex-1 flex-col justify-between py-4">
                        <div>
                            <a href="#" className="mb-1 inline-block text-lg font-bold text-gray-800 transition duration-100 hover:text-gray-500 lg:text-xl">Aave Chat</a>
                        </div>

                        <div>

                            <span className="flex items-center gap-1 text-sm text-gray-500">
                                <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 text-green-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                                </svg>

                                Customized
                            </span>
                        </div>
                    </div>

                    <div className="flex w-full justify-between border-t p-4 sm:w-auto sm:border-none sm:pl-0 lg:p-6 lg:pl-0">
                        <div className="flex flex-col items-start gap-2">
                            <div className="flex h-12 w-32 my-auto overflow-hidden rounded border">
                                <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded inline-flex items-center">
                                    <Link href="/message"><span>Start Chat</span></Link>
                                </button>
                            </div>
                        </div>
                    </div>
                </div>

                <div className="flex flex-wrap gap-x-4 overflow-hidden rounded-lg border sm:gap-y-4 lg:gap-6">
                    <a href="#" className="group relative block h-48 w-32 overflow-hidden bg-gray-100 sm:h-56 sm:w-60">
                        <img src="./uniswap_v4.png" loading="lazy" alt="Photo by Jessica Radanavong" className="h-full w-full object-cover object-center transition duration-200 group-hover:scale-110" />
                    </a>

                    <div className="flex flex-1 flex-col justify-between py-4">
                        <div>
                            <a href="#" className="mb-1 inline-block text-lg font-bold text-gray-800 transition duration-100 hover:text-gray-500 lg:text-xl">Uniswap V4</a>
                        </div>

                        <div>

                            <span className="flex items-center gap-1 text-sm text-gray-500">
                                <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5 text-green-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                                </svg>

                                Customized
                            </span>
                        </div>
                    </div>

                    <div className="flex w-full justify-between border-t p-4 sm:w-auto sm:border-none sm:pl-0 lg:p-6 lg:pl-0">
                        <div className="flex flex-col items-start gap-2">
                            <div className="flex h-12 w-32 my-auto overflow-hidden rounded border">
                                <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded inline-flex items-center">
                                    <span>Subscribe</span>
                                </button>
                            </div>
                        </div>
                    </div>
                </div>

            </div>
        </>
    )
}

export default Marketplace;