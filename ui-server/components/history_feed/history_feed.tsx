import ButtonShowMore from 'components/button_show_more'
import EpisodePreview from 'components/episode_preview'
import parseISO from 'date-fns/parseISO'
import { Episode } from 'models'
import React from 'react'

export interface StateToProps {
  history: Episode[]
  receivedAll: boolean
  isLoadingMore: boolean
}

export interface DispatchToProps {
  loadMore: (offset: number) => void
}

const HistoryFeed: React.FC<StateToProps & DispatchToProps> = ({
  history,
  loadMore,
  receivedAll,
  isLoadingMore,
}) => {
  if (history.length === 0) {
    return (
      <div className="mt-8">
        <div className="text-2xl text-gray-900 tracking-wide">
          {'Your listen history seems empty.'}
        </div>
        <div className="mt-2 text-default text-gray-800 tracking-wide">
          {'Try playing an episode.'}
        </div>
      </div>
    )
  }

  history.sort(
    (a, b) =>
      +parseISO(`${b.lastPlayedAt} +0000`) -
      +parseISO(`${a.lastPlayedAt} +0000`),
  )

  return (
    <div>
      <h1 className="pt-3 pb-1 text-xl tracking-wide font-semibold">
        {'History'}
      </h1>
      <hr className="mb-3" />
      {history.map((episode) => (
        <div key={episode.id} className="mb-6">
          <EpisodePreview episodeId={episode.id} />
        </div>
      ))}
      {!receivedAll && (
        <div className="w-full h-10 mx-auto my-6">
          <ButtonShowMore
            isLoading={isLoadingMore}
            loadMore={() => loadMore(history.length)}
          />
        </div>
      )}
    </div>
  )
}

export default HistoryFeed
