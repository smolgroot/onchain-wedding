import type { FireworksOptions, FireworksTypes } from './types.js';
export declare class Options implements FireworksTypes.Options {
    hue: FireworksTypes.MinMax;
    rocketsPoint: FireworksTypes.MinMax;
    opacity: number;
    acceleration: number;
    friction: number;
    gravity: number;
    particles: number;
    explosion: number;
    mouse: FireworksTypes.Mouse;
    boundaries: FireworksTypes.Boundaries;
    sound: FireworksTypes.Sounds;
    delay: FireworksTypes.MinMax;
    brightness: FireworksTypes.MinMax;
    decay: FireworksTypes.MinMax;
    flickering: number;
    intensity: number;
    traceLength: number;
    traceSpeed: number;
    lineWidth: FireworksTypes.LineWidth;
    lineStyle: FireworksTypes.LineStyle;
    autoresize: boolean;
    constructor();
    update<T extends FireworksOptions>(options: T): void;
}
