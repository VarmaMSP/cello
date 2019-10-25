import rootReducer from 'reducers'
import { applyMiddleware, compose, createStore } from 'redux'
import thunk, { ThunkMiddleware } from 'redux-thunk'
import { AppActions } from 'types/actions'

// NOTE: Do not export this as type
// doing so will make the editor show the entire AppState in suggestions
export interface AppState extends ReturnType<typeof rootReducer> {}

export const makeStore = (initalState?: object) => {
  const composeEnhancers =
    typeof window != 'undefined' &&
    (window as any).__REDUX_DEVTOOLS_EXTENSION_COMPOSE__
      ? (window as any).__REDUX_DEVTOOLS_EXTENSION_COMPOSE__
      : compose

  return createStore(
    rootReducer,
    initalState,
    composeEnhancers(
      applyMiddleware(thunk as ThunkMiddleware<AppState, AppActions>),
    ),
  )
}
