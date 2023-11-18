export {}
declare global{
    namespace NodeJS{
        interface ProcessEnv{
            NEXT_PUBLIC_WORLD_COIN_ID: string
            NEXT_PUBLIC_STORAGE_TOKEN: string
        }
    }
}