import type { Options } from './options.js';
export declare class Resize {
    private readonly options;
    private readonly updateSize;
    private readonly container;
    private resizer;
    constructor(options: Options, updateSize: () => void, container: Element);
    mount(): void;
    unmount(): void;
}
