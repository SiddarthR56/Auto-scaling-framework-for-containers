import os
import pickle
import numpy as np
from keras.models import Sequential
from keras.layers import Bidirectional
from keras.layers import Dropout, LSTM, RepeatVector
from keras.layers import Dense, BatchNormalization
from tensorflow.keras.optimizers import Adam
from sklearn.preprocessing import StandardScaler

Window = 5

class BILSTM:

    def __init__(self) -> None:
        self.model = self.build()
        if os.path.exists('./model.pkl'):
            self.load_model('./model.pkl')
        self.scr = pickle.load(open('./scr.pkl', 'rb'))
        self.scc = pickle.load(open('./scc.pkl', 'rb'))
        self.scm = pickle.load(open('./scm.pkl', 'rb'))


    def build(self):
        model = Sequential()
        model.add(Bidirectional(LSTM(units=64, input_shape=( Window, 3))))
        model.add(RepeatVector(n=Window))
        model.add(Bidirectional(LSTM(units=64)))
        model.add(Dense(15, activation="elu"))
        model.add(Dropout(0.2))
        model.add(Dense(units=1, activation="elu"))

        adamOpt = Adam(lr=0.001, beta_1=0.9, beta_2=0.999, decay=0.0, amsgrad=False)
        model.compile(loss='mean_squared_error', optimizer=adamOpt, metrics=['mae'])

        return model
    
    def load_model(self, path):
        self.model = pickle.load(open(path, 'rb'))

    def normalise_data(self, data):
        cpu = data[:, 0]
        mem = data[:, 2]
        rps = data[:, 1]

        cpu = self.scc.transform(cpu.reshape(-1, 1))
        mem = self.scm.transform(mem.reshape(-1, 1))
        rps = self.scr.transform(rps.reshape(-1, 1))

        return np.concatenate((cpu, rps, mem), axis=1)

    def predict(self, X_test):
        X_test = self.normalise_data(X_test)
        return self.scc.inverse_transform(self.model.predict(np.array([X_test]))[0])
    
    def save_model(self, path):
        pickle.dump(self.model, open(path, 'wb'))

    def train(self, model, X_train, y_train, X_test, y_test):
        model.fit(X_train, y_train,
                batch_size=64,
                epochs=100,
                validation_data=(X_test, y_test))
        return model
    

    
    
