import { enableProdMode, provideZoneChangeDetection } from '@angular/core';
import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';
import { AppModule } from './app/app.module';
import { environment } from './environments/environment';
import dayjs from 'dayjs';
import 'dayjs/locale/zh-cn';

// 全局使用中文 locale，供 dayjs.fromNow() 等相对时间输出使用
dayjs.locale('zh-cn');

if (environment.production) {
    enableProdMode();
}

platformBrowserDynamic()
    .bootstrapModule(AppModule, { applicationProviders: [provideZoneChangeDetection()] })
    .catch((err) => console.error(err));
