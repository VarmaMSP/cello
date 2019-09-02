import { combineReducers, applyMiddleware, createStore, compose } from 'redux'
import thunk, { ThunkMiddleware } from 'redux-thunk'
import entities from '../reducers/entities'
import requests from '../reducers/requests'
import { AppActions } from 'types/actions'

export const rootReducer = combineReducers({
  entities,
  requests,
})

// NOTE: Do not export this as type
// doing so will make the editor to show to entire AppState in suggestions
export interface AppState extends ReturnType<typeof rootReducer> {}

export const makeStore = () => {
  const composeEnhancers =
    typeof window != 'undefined' &&
    (window as any).__REDUX_DEVTOOLS_EXTENSION_COMPOSE__
      ? (window as any).__REDUX_DEVTOOLS_EXTENSION_COMPOSE__
      : compose

  return createStore(
    rootReducer,
    composeEnhancers(
      applyMiddleware(thunk as ThunkMiddleware<AppState, AppActions>),
    ),
  )
}
