import { setTimeout } from 'node:timers/promises';

const task = async (id: number): Promise<number> => {
    console.log(`Task ${id} starting...`);
    await setTimeout(1000);
    console.log(`Task ${id} finished`);
    return id * 2;
};

const main = async () => {
    const numJobs = 10;
    const concurrencyLimit = 3;
    const ids = Array.from({ length: numJobs }, (_, i) => i + 1);
    
    const results: number[] = [];
    const executing: Promise<any>[] = [];

    for (const id of ids) {
        const p = task(id).then((res) => {
            results.push(res);
            executing.splice(executing.indexOf(p), 1);
        });
        executing.push(p);

        if (executing.length >= concurrencyLimit) {
            await Promise.race(executing);
        }
    }

    await Promise.all(executing);
    console.log("All results:", results);
};

main().catch(console.error);