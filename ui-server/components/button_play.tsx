import { beginPlayback } from 'actions/episode'
import classNames from 'classnames'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getEpisodePlayback } from 'selectors/entities/users'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import { Episode, EpisodePlayback } from 'types/app'
import ButtonWithIcon from './button_with_icon'

interface StateToProps {
  playback: EpisodePlayback | undefined
}

interface DispatchToProps {
  playEpisode: (startTime: number) => void
}

interface OwnProps {
  episode: Episode
}

interface Props extends StateToProps, DispatchToProps, OwnProps {}

const ButtonPlay: React.SFC<Props> = ({ playback, playEpisode }) => {
  return (
    <ButtonWithIcon
      className={classNames(
        'flex-none md:w-8 w-6 mx-auto text-gray-600 hover:text-black',
      )}
      icon="play-outline"
      onClick={() => playEpisode(playback ? playback.currentTime : 0)}
    />
  )
}

function mapStateToProps(state: AppState, { episode }: OwnProps): StateToProps {
  return {
    playback: getEpisodePlayback(state, episode.id),
  }
}

function mapDispatchToProps(
  dispatch: Dispatch<AppActions>,
  { episode }: OwnProps,
): DispatchToProps {
  return {
    playEpisode: (startTime: number) =>
      bindActionCreators(beginPlayback, dispatch)(episode.id, startTime),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(ButtonPlay)
