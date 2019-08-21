import React, { Component } from 'react'
import { AudioState } from './components/utils';
import ButtonWithIcon from '../button_with_icon';
import ActionButton from './components/action_button';
import NavbarBottom from '../navbar_bottom/navbar_bottom';
import ProgressBar from './components/progress_bar';

interface Props {
  podcast: string;
  episode: string;
  podcastId: string;
  episodeId: string;
  albumArt: string;
  audioState: AudioState;
  duration: number,
  currentTime: number,
  handleSeek: (t: number) => void;
  handleFastForward: (t: number) => void;
  handleActionButtonPress: () => void;
}

const AudioPlayerMedium: React.SFC<Props> = props => {
  const {
    podcast, episode, audioState,
    currentTime, duration,
    handleSeek, handleActionButtonPress, handleFastForward
  } = props

  return <div className="fixed bottom-0 left-0 flex flex-col justify-between bg-white w-full h-auto px-5 border">
    <div className="flex items-center justify-center w-full h-20">
      <div className="flex-none flex items-center h-full">
        <ButtonWithIcon className="w-10" icon="fast-rewind" onClick={() => handleFastForward(-10)}/>
        <ActionButton audioState={audioState} handleActionButtonPress={handleActionButtonPress}/>
        <ButtonWithIcon className="w-10" icon="fast-forward" onClick={() => handleFastForward(10)}/>
      </div>
      <div className="flex-auto flex flex-col justify-center ml-6 text-center truncate">
        <div className="-mb-1">
          <div className="text-sm font-semibold text-gray-800 leading-tight tracking-wide">{episode}</div>
          <div className="text-sm text-gray-700 leading-loose tracking-wide truncate">{podcast}</div>
        </div>
        <div className="w-full">
          <ProgressBar
            currentTime={currentTime}
            duration={duration}
            handleSeek={handleSeek}
          />
        </div>
      </div>
    </div>
    <div className="mt-2">
      <NavbarBottom/>
    </div>
  </div>
}

export default AudioPlayerMedium
    
