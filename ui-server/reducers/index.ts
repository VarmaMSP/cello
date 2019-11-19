import { combineReducers } from 'redux'
import browser from './browser'
import entities from './entities'
import requests from './requests'
import requests_ from './requests-test'
import ui from './ui'

export default combineReducers({
  ui,
  entities,
  requests,
  browser,
  requests_,
})
