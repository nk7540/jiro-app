import {useNavigation} from '@react-navigation/native';
import React, {useState} from 'react';
import {StyleSheet, View, SafeAreaView} from 'react-native';
import {Button, Input} from 'react-native-elements';
import {ScreenNames} from 'services/navigation';
import {gql, useMutation} from '@apollo/client';
import {CreateUserMutation} from './__generated__/CreateUserMutation';

const CREATE_USER_MUTATION = gql`
  mutation CreateUserMutation($input: CreateUser!) {
    createUser(input: $input)
  }
`;

const SignUp = () => {
  const navigation = useNavigation();
  const [input, setInput] = useState({
    nickname: '',
    email: '',
    password: '',
    passwordConfirmation: '',
  });
  const [createUser] = useMutation<CreateUserMutation>(CREATE_USER_MUTATION, {
    onCompleted: _ => navigation.navigate(ScreenNames.SignIn),
  });

  return (
    <SafeAreaView>
      <View style={styles.form}>
        <Input
          placeholder="ニックネーム"
          value={input.nickname}
          onChangeText={v => setInput({...input, nickname: v})}
        />
        <Input placeholder="email@com" value={input.email} onChangeText={v => setInput({...input, email: v})} />
        <Input
          placeholder="password"
          value={input.password}
          onChangeText={v => setInput({...input, password: v})}
          secureTextEntry
        />
        <Input
          placeholder="password（確認）"
          value={input.passwordConfirmation}
          onChangeText={v => setInput({...input, passwordConfirmation: v})}
          secureTextEntry
        />
        <Button
          title="登録する"
          titleStyle={styles.submitTitle}
          buttonStyle={styles.submitButton}
          onPress={() => createUser({variables: {input: input}})}
        />
      </View>
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  form: {
    alignItems: 'center',
    marginHorizontal: 20,
    marginTop: 50,
  },
  submitButton: {
    paddingHorizontal: 30,
    paddingVertical: 15,
    borderRadius: 20,
    marginBottom: 30,
  },
  submitTitle: {
    fontWeight: 'bold',
    fontSize: 20,
  },
});

export default SignUp;
