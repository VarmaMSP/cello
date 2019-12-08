import { createSelector } from 'reselect'
import { AppState } from 'store'
import { Chart, Podcast } from 'types/app'
import { $Id, MapById, MapOneToMany } from 'types/utilities'

export function getChartById(state: AppState, chartId: string) {
  return state.entities.charts.charts[chartId]
}

export function makeGetCategories() {
  return createSelector<
    AppState,
    MapById<Chart>,
    MapOneToMany<Chart, Chart>,
    Chart[]
  >(
    (state) => state.entities.charts.charts,
    (state) => state.entities.charts.chartTree,
    (all, chartTree) =>
      Object.keys(chartTree)
        .map((id) => all[id])
        .filter((c) => c.type === 'CATEGORY'),
  )
}

export function makeGetSubCategories() {
  return createSelector<
    AppState,
    $Id<Chart>,
    MapById<Chart>,
    $Id<Chart>[],
    Chart[]
  >(
    (state) => state.entities.charts.charts,
    (state, id) => state.entities.charts.chartTree[id],
    (all, ids) => ids.map((id) => all[id]).filter((c) => c.type === 'CATEGORY'),
  )
}

export function makeGetPodcastsInChart() {
  return createSelector<
    AppState,
    $Id<Chart>,
    MapById<Podcast>,
    $Id<Podcast>[],
    Podcast[]
  >(
    (state) => state.entities.podcasts.podcasts,
    (state, chartId) => state.entities.charts.podcastsInCharts[chartId] || [],
    (all, ids) => ids.map((id) => all[id]),
  )
}
