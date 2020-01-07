import { combineReducers } from 'redux'
import entities from './entities'
import history from './history'
import requests from './requests'
import session from './session'
import ui from './ui'
import window from './window'

export default combineReducers({
  ui,
  entities,
  session,
  history,
  window,
  requests,
})
