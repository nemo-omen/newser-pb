import { writable } from "svelte/store";
import type { AuthModel } from "pocketbase";

export const user = writable({} as AuthModel | undefined);