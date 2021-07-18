import database from '@react-native-firebase/database';
import {GeoFire} from 'geofire';
import {Reference} from '@firebase/database-types';

export const firebaseRef = database().ref();
// workaround
// see: https://stackoverflow.com/a/58292506/12309177
export const geoFire = new GeoFire(firebaseRef as unknown as Reference);
