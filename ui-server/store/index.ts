import { applyMiddleware, createStore, compose } from 'redux'
import thunk, { ThunkMiddleware } from 'redux-thunk'
import { AppActions } from 'types/actions'
import rootReducer from '../reducers'

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