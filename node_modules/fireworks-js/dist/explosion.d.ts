import type { FireworksTypes } from './types.js';
export declare class Explosion {
    private x;
    private y;
    private ctx;
    private hue;
    private friction;
    private gravity;
    private flickering;
    private lineWidth;
    private explosionLength;
    private angle;
    private speed;
    private brightness;
    private coordinates;
    private decay;
    private alpha;
    constructor({ x, y, ctx, hue, decay, gravity, friction, brightness, flickering, lineWidth, explosionLength }: FireworksTypes.Explosion);
    update(callback: () => void): void;
    draw(): void;
}
