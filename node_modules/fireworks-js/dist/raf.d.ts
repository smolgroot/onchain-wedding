import type { Options } from './options.js';
export declare class RequestAnimationFrame {
    private readonly options;
    private readonly render;
    tick: number;
    private rafId;
    private fps;
    private tolerance;
    private now;
    constructor(options: Options, render: () => void);
    mount(): void;
    unmount(): void;
}
