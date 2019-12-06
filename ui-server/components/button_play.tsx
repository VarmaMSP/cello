import { startPlayback } from 'actions/playback'
import classNames from 'classnames'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getEpisodeById } from 'selectors/entities/episodes'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import { Episode } from 'types/app'
import ButtonWithIcon from './button_with_icon'

interface StateToProps {
  episode: Episode
}

interface DispatchToProps {
  playEpisode: (startTime: number) => void
}

interface OwnProps {
  episodeId: string
  className: string
}

interface Props extends StateToProps, DispatchToProps, OwnProps {}

const ButtonPlay: React.SFC<Props> = ({ episode, playEpisode, className }) => {
  return (
    <ButtonWithIcon
      className={classNames(
        'flex-none text-indigo-600 hover:text-indigo-700',
        className,
      )}
      icon="play-outline"
      onClick={() => playEpisode((episode.progress * episode.duration) / 100)}
    />
  )
}

function mapStateToProps(
  state: AppState,
  { episodeId }: OwnProps,
): StateToProps {
  return {
    episode: getEpisodeById(state, episodeId),
  }
}

function mapDispatchToProps(
  dispatch: Dispatch<AppActions>,
  { episodeId }: OwnProps,
): DispatchToProps {
  return {
    playEpisode: (startTime: number) =>
      bindActionCreators(startPlayback, dispatch)(episodeId, startTime),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(ButtonPlay)
