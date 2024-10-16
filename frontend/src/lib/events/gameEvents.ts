import { Player } from "csgo-gsi-types";
import { EventsOn } from "../../../wailsjs/runtime/runtime";

type GameEvents =
  | "RoundOver"
  | "GameMap"
  | "GamePhase"
  | "GameRounds"
  | "GameCTScore"
  | "GameTScore"
  | "Player"
  | "RoundPhase"
  | "RoundWinTeam"
  | "BombState"
  | "BombCountDown";

export function onGameEvent(
  event: "RoundOver",
  callback: (data: boolean) => void
): () => void;

export function onGameEvent(
  event: "GameMap",
  callback: (data: string) => void
): () => void;

export function onGameEvent(
  event: "GamePhase",
  callback: (data: string) => void
): () => void;

export function onGameEvent(
  event: "GameRounds",
  callback: (data: number) => void
): () => void;

export function onGameEvent(
  event: "GameCTScore",
  callback: (data: number) => void
): () => void;

export function onGameEvent(
  event: "GameTScore",
  callback: (data: number) => void
): () => void;

export function onGameEvent(
  event: "Player",
  callback: (data: Player) => void
): () => void;

export function onGameEvent(
  event: "RoundPhase",
  callback: (data: string) => void
): () => void;

export function onGameEvent(
  event: "RoundWinTeam",
  callback: (data: string) => void
): () => void;

export function onGameEvent(
  event: "BombState",
  callback: (data: string) => void
): () => void;

export function onGameEvent(
  event: "BombCountDown",
  callback: (data: number) => void
): () => void;

export function onGameEvent(
  event: GameEvents,
  callback: (...data: any) => void
): () => void {
  return EventsOn(event, callback);
}
