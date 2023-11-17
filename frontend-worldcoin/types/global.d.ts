export {}
declare global{
    namespace NodeJS{
        interface ProcessEnv{
            NEXT_PUBLIC_WORLD_COIN_ID: string
        }
    }
}