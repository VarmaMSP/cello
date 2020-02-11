import { Epic, ofType } from 'redux-observable'
import { from, Observable, of } from 'rxjs'
import {
  catchError,
  concatMap,
  debounceTime,
  filter,
  map,
  throttleTime,
} from 'rxjs/operators'
import { AppState } from 'store'
import * as T from 'types/actions'
import { UpdateTextAction } from 'types/actions/ui/search_bar'
import { doFetch } from 'utils/fetch'

const searchEpic: Epic<T.AppActions, T.AppActions, AppState> = (action$) =>
  action$.pipe(
    ofType(T.SEARCH_BAR_UPDATE_TEXT),
    filter<UpdateTextAction>(({ text }) => text.trim().length > 0),
    debounceTime<UpdateTextAction>(300),
    concatMap<UpdateTextAction, Observable<T.AppActions>>((action) =>
      from(
        doFetch({
          method: 'GET',
          urlPath: `/suggest?query=${action.text}`,
        }),
      ).pipe(
        map(({ searchSuggestions }) => ({
          type: T.SEARCH_BAR_UPDATE_SEARCH_SUGGESTIONS,
          suggestions: searchSuggestions,
        })),
        catchError<any, Observable<T.AppActions>>(() =>
          of({ type: 'CONTINUE' }),
        ),
      ),
    ),
    throttleTime<T.AppActions>(100),
  )

export default searchEpic
