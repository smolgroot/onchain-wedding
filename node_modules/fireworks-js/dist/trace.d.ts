import type { FireworksTypes } from './types.js';
export declare class Trace {
    private x;
    private y;
    private sx;
    private sy;
    private dx;
    private dy;
    private ctx;
    private hue;
    private speed;
    private acceleration;
    private traceLength;
    private totalDistance;
    private angle;
    private brightness;
    private coordinates;
    private currentDistance;
    constructor({ x, y, dx, dy, ctx, hue, speed, traceLength, acceleration }: FireworksTypes.Trace);
    update(callback: (x: number, y: number, hue: number) => void): void;
    draw(): void;
}
