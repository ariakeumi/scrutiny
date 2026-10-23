import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { FileSizePipe } from './file-size.pipe';
import { DeviceSortPipe } from './device-sort.pipe';
import { TemperaturePipe } from './temperature.pipe';
import { DeviceTitlePipe } from './device-title.pipe';
import { DeviceStatusPipe } from './device-status.pipe';
import { DeviceStatusLabelPipe } from './device-status-label.pipe';
import { DeviceHoursPipe } from './device-hours.pipe';
import { LatencyPipe } from './latency.pipe';

@NgModule({
    imports: [CommonModule, FormsModule, ReactiveFormsModule, FileSizePipe, DeviceSortPipe, TemperaturePipe, DeviceTitlePipe, DeviceStatusPipe, DeviceStatusLabelPipe, DeviceHoursPipe, LatencyPipe],
    exports: [CommonModule, FormsModule, ReactiveFormsModule, FileSizePipe, DeviceSortPipe, DeviceTitlePipe, DeviceStatusPipe, DeviceStatusLabelPipe, TemperaturePipe, DeviceHoursPipe, LatencyPipe],
})
export class SharedModule {}
