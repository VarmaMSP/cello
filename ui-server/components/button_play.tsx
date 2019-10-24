import { beginPlayback } from 'actions/episode'
import classNames from 'classnames'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getCurrentUserPlayback } from 'selectors/entities/episodes'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import { EpisodePlayback } from 'types/app'
import ButtonWithIcon from './button_with_icon'

interface StateToProps {
  playback?: EpisodePlayback
}

interface DispatchToProps {
  playEpisode: (startTime: number) => void
}

interface OwnProps {
  episodeId: string
  className: string
}

interface Props extends StateToProps, DispatchToProps, OwnProps {}

const ButtonPlay: React.SFC<Props> = ({ playback, playEpisode, className }) => {
  return (
    <ButtonWithIcon
      className={classNames(
        'flex-none text-gray-600 hover:text-black',
        className,
      )}
      icon="play-outline"
      onClick={() => playEpisode(playback ? playback.currentTime : 0)}
    />
  )
}

function mapStateToProps(
  state: AppState,
  { episodeId }: OwnProps,
): StateToProps {
  return {
    playback: getCurrentUserPlayback(state, episodeId),
  }
}

function mapDispatchToProps(
  dispatch: Dispatch<AppActions>,
  { episodeId }: OwnProps,
): DispatchToProps {
  return {
    playEpisode: (startTime: number) =>
      bindActionCreators(beginPlayback, dispatch)(episodeId, startTime),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(ButtonPlay)
