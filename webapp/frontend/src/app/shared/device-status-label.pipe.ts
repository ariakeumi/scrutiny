import { Pipe, PipeTransform } from '@angular/core';

const STATUS_LABELS: { [key: string]: string } = {
    passed: '安全',
    failed: '警告',
    unknown: '未知',
    warn: '警告',
};

/**
 * Maps a device status string ('passed' | 'failed' | 'unknown') to a
 * Chinese display label. Used for display only; status comparisons in
 * templates/code should keep using the raw English keys.
 */
@Pipe({ name: 'deviceStatusLabel' })
export class DeviceStatusLabelPipe implements PipeTransform {
    transform(status: string | null | undefined): string {
        if (!status) {
            return '';
        }
        // strip any "failed: <reason>" suffix down to the base status
        const base = status.split(':')[0].trim();
        return STATUS_LABELS[base] || status;
    }
}
