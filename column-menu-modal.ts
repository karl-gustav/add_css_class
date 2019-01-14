import {
    Component,
    Input,
    Output,
    EventEmitter,
    ChangeDetectionStrategy,
    OnInit
} from '@angular/core';
import {IModalOptions, IUniModal} from '../../uni-modal/interfaces';
import {UniTableConfig} from '../unitable/config/unitableConfig';
import {UniTableColumn} from '../unitable/config/unitableColumn';

import * as _ from 'lodash';

@Component({
    selector: 'column-menu-modal',
    template: `
        <section role="dialog" class="uni-modal" [ngClass]="{'advanced': tableConfig?.advancedColumnMenu}">
            <header>
                <h1>{{'Kolonneoppsett'}}</h1>
            </header>
            <article class="scrollable">
                <p>Her kan du bestemme synlighet, tittel, rekkefølge og posisjon på kolonnene.</p>
                <p>"Hopp til kolonne" definere hvilken kolonne man skal gå til ved tab/enter klikk.</p>
                <p>For å endre posisjon på en kolonne drar du <i class="material-icons move-icon">menu</i> ikonet opp eller ned.</p>
                <table>
                    <thead>
                        <tr>
                            <th class="visibility-col">Synlig</th>
                            <th class="title-col">Tittel</th>
                            <th class="jump-col">Hopp til kolonne</th>
                            <ng-container *ngIf="tableConfig?.advancedColumnMenu">
                                <th>Felt eller spørring</th>
                                <th>Summeringsfunksjon</th>
                                <th>Alias</th>
                            </ng-container>
                            <th><!-- columnMenu toggle icon --></th>
                        </tr>
                    </thead>

                    <tbody>
                        <tr *ngFor="let column of columns; let idx = index"
                            draggable="true"
                            (dragstart)="onDragStart($event, idx)"
                            (dragover)="onDragOver($event)"
                            (dragleave)="onDragLeave($event)"
                            (drop)="onDrop($event, idx)"
                            (dragend)="onDragEnd($event)">

                            <td>
                                <input type="checkbox"
                                    [checked]="column.visible"
                                    (change)="visibilityChanged(idx, column)"
                                />
                                <label class="checkbox-label" (click)="visibilityChanged(idx, column)">Synlig</label>
                            </td>

                            <td>
                                {{column.header}}
                            </td>

                            <td>
                                <select [value]="column.jumpToColumn || ''"
                                        (change)="inputChange($event, 'jumpToColumn', column, idx)">
                                    <option value=""></option>
                                    <option *ngFor="let col of columns"
                                            value="{{col.field}}">
                                        {{col.header}}
                                    </option>
                                </select>
                            </td>

                            <ng-container *ngIf="tableConfig?.advancedColumnMenu">
                                <td>
                                    <input type="text"
                                        [value]="column.field"
                                        (change)="inputChange($event, 'field', column, idx)"
                                        placeholder="Feltnavn eller formel"
                                    />
                                </td>

                                <td>
                                    <select [value]="column.sumFunction || ''"
                                            (change)="inputChange($event, 'sumFunction', column, idx)">
                                        <option value=""></option>
                                        <option value="sum">Sum</option>
                                        <option value="avg">Gjennomsnitt</option>
                                        <option value="min">Laveste</option>
                                        <option value="max">Høyeste</option>
                                    </select>
                                </td>

                                <td>
                                    <input type="text"
                                        [value]="column.alias || ''"
                                        (change)="inputChange($event, 'alias', column, idx)"
                                    />
                                </td>
                            </ng-container>

                            <td>
                                <i class="material-icons move-icon">menu</i>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </article>

            <footer>
                <button type="submit" class="good" (click)="close(true)">Lagre</button>
                <button (click)="resetAll()">Nullstill</button>
                <button class="cancel" (click)="close(false)">Avbryt</button>
            </footer>
        </section>

                <select-products-for-bulk-access
                    *ngIf="currentPage === PageType.selectProducts"
                    [data]="grantAccessData"
                    (stepComplete)="productsSelected = $event"
                ></select-products-for-bulk-access>
