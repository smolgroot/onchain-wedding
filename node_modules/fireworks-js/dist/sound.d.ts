import type { Options } from './options.js';
declare global {
    interface Window {
        webkitAudioContext: typeof AudioContext;
    }
}
export declare class Sound {
    private readonly options;
    private buffers;
    private audioContext;
    private onInit;
    constructor(options: Options);
    private get isEnabled();
    private get soundOptions();
    private init;
    private loadSounds;
    play(): void;
}
