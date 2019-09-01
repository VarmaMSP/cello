import { Dispatch, bindActionCreators } from 'redux'
import { getPodcast } from '../../actions/podcast'
import PodcastPage from './podcast'
import { connect } from 'react-redux'

function mapDispatchToProps(dispatch: Dispatch) {
  return {
    getPodcast: bindActionCreators(getPodcast, dispatch),
  }
}

export default connect(
  null,
  mapDispatchToProps,
)(PodcastPage)
