import { combineReducers, Reducer } from 'redux'
import * as T from 'types/actions'
import { Chart } from 'types/app'

const charts: Reducer<{ [id: string]: Chart }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.RECEIVED_PODCAST_CHART:
      return { ...state, [action.chart.id]: action.chart }
    case T.RECEIVED_PODCAST_CHARTS:
      return {
        ...state,
        ...action.charts.reduce<{ [id: string]: Chart }>(
          (acc, c) => ({ ...acc, [c.id]: c }),
          {},
        ),
      }
    default:
      return state
  }
}

const chartTree: Reducer<{ [id: string]: string[] }, T.AppActions> = (
  state = {},
  action,
) => {
  switch (action.type) {
    case T.RECEIVED_PODCAST_CHART:
      return !!action.chart.parentId
        ? {
            ...state,
            [action.chart.parentId]: [
              ...new Set([
                ...(state[action.chart.parentId] || []),
                action.chart.id,
              ]),
            ],
          }
        : state
    case T.RECEIVED_PODCAST_CHARTS:
      return {
        ...state,
        ...action.charts.reduce<{ [id: string]: string[] }>(
          (acc, c) =>
            !!c.parentId
              ? { ...acc, [c.parentId]: [...(acc[c.parentId] || []), c.id] }
              : acc,
          {},
        ),
      }
    default:
      return state
  }
}

const podcastsInCharts: Reducer<
  { [chartId: string]: string[] },
  T.AppActions
> = (state = {}, action) => {
  switch (action.type) {
    case T.RECEIVED_CHART_PODCASTS:
      return { ...state, [action.chartId]: action.podcasts.map((p) => p.id) }
    default:
      return state
  }
}

export default combineReducers({
  charts,
  chartTree,
  podcastsInCharts,
})
