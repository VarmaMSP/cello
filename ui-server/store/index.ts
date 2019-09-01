import { combineReducers, applyMiddleware, createStore } from 'redux'
import thunk, { ThunkMiddleware } from 'redux-thunk'
import entities from '../reducers/entities'
import requests from '../reducers/requests'
import { AppActions } from 'types/actions'

export const rootReducer = combineReducers({
  entities,
  requests,
})

export type AppState = ReturnType<typeof rootReducer>

export const store = createStore(
  rootReducer,
  applyMiddleware(thunk as ThunkMiddleware<AppState, AppActions>),
)
