import { connect } from 'react-redux'
import {
  getCategoryById,
  makeGetPodcastsInList,
} from 'selectors/entities/podcast_lists'
import { AppState } from 'store'
import PodcastList, { OwnProps, StateToProps } from './podcasts_list'

function makeMapStateToProps() {
  const getPodcastsInList = makeGetPodcastsInList()

  return (state: AppState, { listId }: OwnProps): StateToProps => {
    return {
      list: getCategoryById(state, listId),
      podcasts: getPodcastsInList(state, listId),
    }
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(
  makeMapStateToProps(),
)(PodcastList)
