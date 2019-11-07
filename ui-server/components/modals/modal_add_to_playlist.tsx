import { createPlaylist } from 'actions/playlist'
import React from 'react'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import { AppState } from 'store'
import { AppActions } from 'types/actions'
import { Playlist } from 'types/app'
import ModalContainer from './components/modal_container'
import Overlay from './components/overlay'

interface StateToProps {
  playlists: Playlist[]
}

interface DispatchToProps {
  createPlaylist: (
    title: string,
    privacy: 'PUBLIC' | 'PRIVATE' | 'ANONYMOUS',
  ) => void
}

interface OwnProps {
  closeModal: () => void
}

interface Props extends StateToProps, DispatchToProps, OwnProps {}

const ModalAddToPlaylist: React.SFC<Props> = (props) => {
  return (
    <Overlay background="rgba(0, 0, 0, 0.75)">
      <ModalContainer
        handleClose={props.closeModal}
        closeUponClicking="OVERLAY"
      >
        <div className="flex h-full">
          <div className="w-3/5 h-full border-r">Add to playlist</div>
          <div className="w-2/5">
            <button
              className="w-32 h-12 bg-blue"
              onClick={() => props.createPlaylist('NEW PLAYLIST', 'PUBLIC')}
            >
              {'Create Playlist'}
            </button>
          </div>
        </div>
      </ModalContainer>
    </Overlay>
  )
}

function mapStateToProps(_: AppState): StateToProps {
  return {
    playlists: [],
  }
}

function mapDispatchToProps(dispatch: Dispatch<AppActions>): DispatchToProps {
  return {
    createPlaylist: (
      title: string,
      privacy: 'PUBLIC' | 'PRIVATE' | 'ANONYMOUS',
    ) => bindActionCreators(createPlaylist, dispatch)(title, privacy),
  }
}

export default connect<StateToProps, DispatchToProps, OwnProps, AppState>(
  mapStateToProps,
  mapDispatchToProps,
)(ModalAddToPlaylist)
