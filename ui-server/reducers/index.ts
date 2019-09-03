import ui from './ui'
import entities from './entities'
import requests from './requests'
import { combineReducers } from 'redux'

export default combineReducers({
  ui,
  entities,
  requests,
})
