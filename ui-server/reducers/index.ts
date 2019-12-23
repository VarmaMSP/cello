import { combineReducers } from 'redux'
import browser from './browser'
import entities from './entities'
import requests from './requests'
import session from './session'
import ui from './ui'

export default combineReducers({
  ui,
  entities,
  session,
  browser,
  requests,
})
