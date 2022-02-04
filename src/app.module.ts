import { UsersModule } from './modules/users/users.module'
import { AuthModule } from './modules/auth/auth.module'
import { NotesModule } from './modules/notes/notes.module'
import { Module } from '@nestjs/common'

@Module({
  imports: [UsersModule, AuthModule, NotesModule],
  controllers: [],
  providers: [],
})
export class AppModule {}
