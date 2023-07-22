export declare function floor(num: number): number;
export declare function randomFloat(min: number, max: number): number;
export declare function randomInt(min: number, max: number): number;
export declare function getDistance(x: number, y: number, dx: number, dy: number): number;
export declare function hsla(hue: number, lightness: number, alpha?: number): string;
interface IObject {
    [key: string]: any;
    length?: never;
}
declare type IUnionToIntersection<U> = (U extends any ? (k: U) => void : never) extends (k: infer I) => void ? I : never;
export declare const deepMerge: <T extends IObject[]>(...objects: T) => IUnionToIntersection<T[number]>;
export declare function debounce<T extends (...args: any[]) => void>(fn: T, ms: number): (...args: Parameters<T>) => void;
export {};
