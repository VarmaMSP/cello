import classnames from 'classnames'
import { iconMap } from 'components/icon'
import { SearchSuggestion } from 'models'
import React, { useEffect, useState } from 'react'
import { getImageUrl, stopEventPropagation } from 'utils/dom'

export interface StateToProps {
  suggestions: SearchSuggestion[]
}

export interface DispatchToProps {
  loadResultsPage: (text: string) => void
  loadPodcastPage: (podcastUrlParam: string) => void
}

const SearchSuggestionsList: React.FC<StateToProps & DispatchToProps> = ({
  suggestions,
  loadResultsPage,
  loadPodcastPage,
}) => {
  const [cursor, setCursor] = useState<number>(0)

  const handleOnClick = (s: SearchSuggestion) => () => {
    SearchSuggestion.isPodcast(s)
      ? loadPodcastPage(s.i)
      : loadResultsPage(s.header.replace(/<\/?em>/g, ''))
  }

  const handleOnKeyDown = (e: any) => {
    if (e.keyCode === 38 && cursor > 0) {
      setCursor(cursor - 1)
    } else if (e.keyCode === 40 && cursor < suggestions.length - 1) {
      setCursor(cursor + 1)
    } else if (e.keyCode === 13) {
      SearchSuggestion.isPodcast(suggestions[cursor])
        ? loadPodcastPage(suggestions[cursor].i)
        : loadResultsPage(suggestions[cursor].header.replace(/<\/?em>/g, ''))
    }
  }

  useEffect(() => {
    cursor >= suggestions.length && setCursor(0)
  }, [suggestions.length])

  useEffect(() => {
    document.addEventListener('keydown', handleOnKeyDown)
    return () => document.removeEventListener('keydown', handleOnKeyDown)
  }, [cursor, suggestions.map((x) => x.id).join()])

  return (
    <div
      style={{ width: '32rem' }}
      className="z-10 px-2 py-2 bg-white border border-blue-400 rounded-lg"
    >
      {suggestions.map((s, i) =>
        SearchSuggestion.isPodcast(s)
          ? renderItemSuggestion(s, handleOnClick(s), cursor === i)
          : renderTextSuggestion(s, handleOnClick(s), cursor === i),
      )}
    </div>
  )
}

function renderTextSuggestion(
  s: SearchSuggestion,
  onClick: () => void,
  active: boolean,
): JSX.Element {
  const Icon = iconMap[s.i === 'S' ? 'search' : 'enter']

  return (
    <div
      key={s.id}
      className={classnames(
        'search-suggestion flex items-center px-3 py-1 hover:bg-gray-200 cursor-pointer rounded',
        { 'bg-gray-200': active },
      )}
      onClick={onClick}
      onPointerDown={stopEventPropagation}
      onMouseDown={stopEventPropagation}
      onTouchStart={stopEventPropagation}
    >
      <Icon className="flex-none w-4 h-4 mr-4 fill-current text-gray-800" />
      <div
        className="lowercase text-base text-gray-900 leading-loose line-clamp-1"
        dangerouslySetInnerHTML={{ __html: s.header }}
      />
    </div>
  )
}

function renderItemSuggestion(
  s: SearchSuggestion,
  onClick: () => void,
  active: boolean,
): JSX.Element {
  return (
    <div
      key={s.id}
      className={classnames(
        'search-suggestion flex p-3 hover:bg-gray-200 cursor-pointer rounded',
        { 'bg-gray-200': active },
      )}
      onClick={onClick}
      onPointerDown={stopEventPropagation}
      onMouseDown={stopEventPropagation}
      onTouchStart={stopEventPropagation}
    >
      <img
        src={getImageUrl(s.i)}
        className="flex-none w-12 h-12 mr-3 border rounded"
      />
      <div>
        <div
          className="text-base text-gray-900 line-clamp-1"
          dangerouslySetInnerHTML={{ __html: s.header }}
        />
        <div
          className="text-sm text-gray-900 line-clamp-1"
          dangerouslySetInnerHTML={{ __html: s.subHeader }}
        />
      </div>
    </div>
  )
}

export default SearchSuggestionsList