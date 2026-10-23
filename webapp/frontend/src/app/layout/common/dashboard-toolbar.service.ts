import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

export type DashboardToolbarAction = 'toggleArchived' | 'openSettings' | 'runCollectors';

/**
 * 全局顶栏与 dashboard 页面之间共享的操作通道。
 * 顶栏的图标按钮通过 dispatch() 把用户操作转发给当前 dashboard 组件执行；
 * dashboard 组件反过来把 showArchived / isTriggering 状态同步回顶栏用于高亮。
 */
@Injectable({ providedIn: 'root' })
export class DashboardToolbarService {
    readonly actions$ = new Subject<DashboardToolbarAction>();

    showArchived = false;
    isTriggering = false;

    dispatch(action: DashboardToolbarAction): void {
        this.actions$.next(action);
    }
}
