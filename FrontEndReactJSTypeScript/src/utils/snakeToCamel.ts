// Function to convert snake_case to camelCase
export function snakeToCamel(obj: any): any {
    if (Array.isArray(obj)) {
        return obj.map(item => snakeToCamel(item));
    } else if (obj !== null && typeof obj === 'object') {
        return Object.entries(obj).reduce((acc, [key, value]) => {
            const camelKey = key.replace(/_./g, match => match[1].toUpperCase());
            acc[camelKey] = snakeToCamel(value);
            return acc;
        }, {} as any);
    }
    return obj;
}