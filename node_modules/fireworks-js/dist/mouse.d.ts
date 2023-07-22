import type { Options } from './options.js';
export declare class Mouse {
    private readonly options;
    private readonly canvas;
    active: boolean;
    x: number;
    y: number;
    constructor(options: Options, canvas: HTMLCanvasElement);
    private get mouseOptions();
    mount(): void;
    unmount(): void;
    private usePointer;
    private pointerDown;
    private pointerUp;
    private pointerMove;
}
