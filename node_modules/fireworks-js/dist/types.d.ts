import type { Fireworks } from './fireworks.js';
export declare namespace FireworksTypes {
    interface Options {
        hue: MinMax;
        rocketsPoint: MinMax;
        opacity: number;
        acceleration: number;
        friction: number;
        gravity: number;
        particles: number;
        explosion: number;
        mouse: Mouse;
        boundaries: Boundaries;
        sound: Sounds;
        delay: MinMax;
        brightness: MinMax;
        decay: MinMax;
        flickering: number;
        intensity: number;
        traceLength: number;
        traceSpeed: number;
        lineWidth: LineWidth;
        lineStyle: LineStyle;
        autoresize: boolean;
    }
    type LineStyle = 'round' | 'square';
    interface Mouse {
        click: boolean;
        move: boolean;
        max: number;
    }
    interface Boundaries {
        x: number;
        y: number;
        width: number;
        height: number;
        debug: boolean;
    }
    interface Sounds {
        enabled: boolean;
        files: string[];
        volume: MinMax;
    }
    interface LineWidth {
        explosion: MinMax;
        trace: MinMax;
    }
    interface MinMax {
        min: number;
        max: number;
    }
    interface Sizes {
        width: number;
        height: number;
    }
    interface Trace {
        x: number;
        y: number;
        dx: number;
        dy: number;
        ctx: CanvasRenderingContext2D;
        hue: number;
        speed: number;
        acceleration: number;
        traceLength: number;
    }
    interface Explosion {
        x: number;
        y: number;
        ctx: CanvasRenderingContext2D;
        hue: number;
        friction: number;
        gravity: number;
        explosionLength: number;
        flickering: boolean;
        lineWidth: number;
        brightness: MinMax;
        decay: MinMax;
    }
}
export declare type FireworksOptions = RecursivePartial<FireworksTypes.Options>;
export interface FireworksHandlers extends Pick<Fireworks, 'isRunning' | 'start' | 'launch' | 'pause' | 'clear' | 'updateOptions' | 'updateBoundaries' | 'updateSize' | 'currentOptions'> {
    waitStop(): Promise<void>;
    stop(): void;
}
export declare type RecursivePartial<T> = {
    [P in keyof T]?: T[P] extends (infer U)[] ? RecursivePartial<U>[] : T[P] extends object ? RecursivePartial<T[P]> : T[P];
};
