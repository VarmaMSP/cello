import { getPlaylistPageData } from 'actions/playlist'
import ButtonSignin from 'components/button_signin'
import { iconMap } from 'components/icon'
import PlaylistView from 'components/playlist_view/playlist_view'
import React from 'react'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { getPlaylistPageStatus } from 'selectors/request'
import { getIsUserSignedIn } from 'selectors/session'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import { PageContext } from 'types/utilities'

interface StateToProps {
  isUserSignedIn: boolean
  isLoading: boolean
}

interface DispatchToProps {
  loadPageData: () => void
}

interface OwnProps {
  scrollY: number
}

type Props = StateToProps & DispatchToProps & OwnProps

class PlaylistPage extends React.Component<Props> {
  static async getInitialProps(ctx: PageContext): Promise<void> {
    const { store } = ctx

    if (getIsUserSignedIn(store.getState())) {
      await bindActionCreators(getPlaylistPageData, store.dispatch)()
    }
  }

  componentDidUpdate(prevProps: Props) {
    const { isUserSignedIn } = this.props
    if (isUserSignedIn && !prevProps.isUserSignedIn) {
      this.props.loadPageData()
    }
  }

  componentDidMount() {
    window.window.scrollTo(0, this.props.scrollY)
  }

  render() {
    const { isUserSignedIn, isLoading } = this.props

    if (!isUserSignedIn || isLoading) {
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

    return (
      <div>
        <PlaylistView />
      </div>
    )
  }
}

function mapStateToProps(state: AppState): StateToProps {
  return {
    isUserSignedIn: getIsUserSignedIn(state),
    isLoading: getPlaylistPageStatus(state) === 'IN_PROGRESS',
  }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    loadPageData: bindActionCreators(getPlaylistPageData, dispatch),
  }
}

export default connect<StateToProps, {}, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(PlaylistPage)
