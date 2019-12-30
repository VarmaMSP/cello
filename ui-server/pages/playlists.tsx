import { getPlaylist, getPlaylistLibrary } from 'actions/playlist'
import ButtonSignin from 'components/button_signin'
import { iconMap } from 'components/icon'
import PageLayout from 'components/page_layout'
import PlaylistsLibrary from 'components/playlist_library'
import PlaylistView from 'components/playlist_view'
import React from 'react'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getIsUserSignedIn } from 'selectors/session'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import { PageContext } from 'types/utilities'
import { getIdFromUrlParam } from 'utils/utils'

interface StateToProps {
  isUserSignedIn: boolean
}

interface DispatchToProps {
  loadPlaylistLibrary: () => void
}

interface OwnProps {
  playlistUrlParam?: string
  activeTab?: string
  skipLoad?: boolean
  scrollY: number
}

type Props = StateToProps & DispatchToProps & OwnProps

class PlaylistPage extends React.Component<Props> {
  static async getInitialProps({ query, store }: PageContext): Promise<void> {
    // load for playlist page
    if (!!query['playlistUrlParam']) {
      await bindActionCreators(
        getPlaylist,
        store.dispatch,
      )(query['playlistUrlParam'] as string)
      return
    }

    // load data for playlist library page
    if (getIsUserSignedIn(store.getState())) {
      await bindActionCreators(getPlaylistLibrary, store.dispatch)()
    }
  }

  componentDidUpdate(prevProps: Props) {
    const { isUserSignedIn, playlistUrlParam } = this.props
    if (!!!playlistUrlParam && isUserSignedIn && !prevProps.isUserSignedIn) {
      this.props.loadPlaylistLibrary()
    }
  }

  componentDidMount() {
    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    const { isUserSignedIn, playlistUrlParam, activeTab } = this.props

    // Playlist Page
    if (!!playlistUrlParam) {
      const playlistId = getIdFromUrlParam(playlistUrlParam)
      return (
        <PageLayout>
          <PlaylistView playlistId={playlistId} activeTab={activeTab} />
          <div />
        </PageLayout>
      )
    }

    if (!isUserSignedIn) {
      const PlaylistIcon = iconMap['playlist']
      return (
        <>
          <div className="mx-auto mt-24">
            <PlaylistIcon className="w-12 h-12 mx-auto fill-current text-gray-700" />
            <h1 className="text-center text-xl text-gray-700 my-6">
              {'Sign in to create playlists'}
            </h1>
            <div className="w-32 mx-auto">
              <ButtonSignin />
            </div>
          </div>
        </>
      )
    }

    // Playlist Library Page
    return (
      <PageLayout>
        <PlaylistsLibrary />
        <div />
      </PageLayout>
    )
  }
}

function mapStateToProps(state: AppState): StateToProps {
  return {
    isUserSignedIn: getIsUserSignedIn(state),
  }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    loadPlaylistLibrary: bindActionCreators(getPlaylistLibrary, dispatch),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(PlaylistPage)
