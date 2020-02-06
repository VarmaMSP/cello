import searchEpic from 'epics/searchEpic';
import rootReducer from 'reducers';
import { applyMiddleware, compose, createStore } from 'redux';
import { createEpicMiddleware } from 'redux-observable';
import thunk, { ThunkMiddleware } from 'redux-thunk';
import { AppActions } from 'types/actions';

// NOTE: Do not export this as type
// doing so will make the editor show the entire AppState in suggestions
export interface AppState extends ReturnType<typeof rootReducer> {}

const epicMiddleware = createEpicMiddleware<AppActions, AppActions, AppState>();

export const makeStore = (initalState?: AppState) => {
  const composeEnhancers =
    typeof window != 'undefined' &&
    (window as any).__REDUX_DEVTOOLS_EXTENSION_COMPOSE__
      ? (window as any).__REDUX_DEVTOOLS_EXTENSION_COMPOSE__
      : compose

  const store =  createStore(
    rootReducer,
    initalState,
    composeEnhancers(
      applyMiddleware(thunk as ThunkMiddleware<AppState, AppActions>, epicMiddleware),
    ),
  )

  epicMiddleware.run(searchEpic)

  return store
}
